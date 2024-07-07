package internal

import (
	"encoding/json"
	"fmt"
)

type Metric struct {
	// Success rate of requests
	Success float64 `json:"success"`
	// Total number of requests made
	Request int `json:"request"`
	// Map of response codes
	ResponseCodes map[string]int `json:"response_codes"`
	// Collection of errors on attacks
	Errors []string `json:"errors"`
}

func NewMetric() *Metric {
	return &Metric{
		Success:       0,
		Request:       0,
		ResponseCodes: make(map[string]int),
		Errors:        make([]string, 0),
	}
}

func (metric *Metric) init() {
	if metric.Errors == nil {
		metric.Errors = make([]string, 0)
	}

	if metric.ResponseCodes == nil {
		metric.ResponseCodes = make(map[string]int)
	}
}

func (metric *Metric) Add(result *FilteredResult) {
	metric.init()
	metric.Request++

	if result.ResponseCode != "" {
		_, ok := metric.ResponseCodes[result.ResponseCode]
		if !ok {
			metric.ResponseCodes[result.ResponseCode] = 1
		} else {
			metric.ResponseCodes[result.ResponseCode]++
		}
	}

	if result.ResponseCode == "000" {
		metric.Success++
	}

	if result.Error != nil {
		metric.Errors = append(metric.Errors, result.Error.Error())
	}
}

func (metric *Metric) Close() {
	metric.Success = (metric.Success / float64(metric.Request)) * 100
}

func (metric *Metric) ShowResult() error {
	// By default, printing result into a json
	jsonByte, err := json.Marshal(metric)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonByte))
	return nil
}
