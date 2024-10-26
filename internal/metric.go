package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/influxdata/tdigest"
	"github.com/yurichandra/ashito/helper"
)

type Metric struct {
	// Earliest represent the timestamp of first attack request
	Earliest time.Time `json:"earliest"`
	// Latest represent the timestamp of the last attack request
	Latest time.Time `json:"latest"`
	// Duration represents total duration of attack
	Duration time.Duration `json:"duration"`
	// Rate represent request rate to attack
	Rate float64 `json:"rate"`
	// Throughput represent number of success request divided by total duration in second
	Throughput float64 `json:"throughput"`
	// Latency represent latency object consist of mean, P50, P90, P95 and P99
	Latency *Latency `json:"latency"`
	// Success rate of requests
	Success float64 `json:"success"`
	// Total number of requests made
	Request int `json:"request"`
	// Map of response codes
	ResponseCodes map[string]int `json:"response_codes"`
	// Collection of errors on attacks
	Errors []string `json:"errors"`
}

type MetricOutput struct {
	// Earliest represent the timestamp of first attack request
	Earliest time.Time `json:"earliest"`
	// Latest represent the timestamp of the last attack request
	Latest time.Time `json:"latest"`
	// Duration represents total duration of attack
	Duration string `json:"duration"`
	// Rate represent request rate to attack
	Rate float64 `json:"rate"`
	// Throughput represent number of success request divided by total duration in second
	Throughput float64 `json:"throughput"`
	// Latency represent latency object consist of mean, P50, P90, P95 and P99
	Latency *LatencyOutput `json:"latency"`
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

	if metric.Latency == nil {
		metric.Latency = &Latency{
			estimator: tdigest.NewWithCompression(100),
		}
	}
}

func (metric *Metric) Add(result *FilteredResult) {
	metric.init()
	metric.Request++
	metric.Latency.Add(result.Latency)
	metric.Latest = result.ExecutedAt

	if metric.Earliest.IsZero() {
		metric.Earliest = result.ExecutedAt
	}

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
	metric.Latency.P50 = time.Duration(metric.Latency.estimator.Quantile(0.50))
	metric.Latency.P90 = time.Duration(metric.Latency.estimator.Quantile(0.90))
	metric.Latency.P95 = time.Duration(metric.Latency.estimator.Quantile(0.95))
	metric.Latency.P99 = time.Duration(metric.Latency.estimator.Quantile(0.99))

	metric.Duration = metric.Latest.Sub(metric.Earliest)

	if sec := metric.Duration.Seconds(); sec > 0 {
		metric.Rate = float64(metric.Request) / sec
		metric.Throughput = metric.Success / sec
	}
}

func (metric *Metric) ShowResult() error {
	// By default, printing result into a json
	metricOutput := MetricOutput{
		Earliest:   metric.Earliest,
		Latest:     metric.Latest,
		Duration:   helper.FormatDuration(metric.Duration),
		Rate:       metric.Rate,
		Throughput: metric.Throughput,
		Latency: &LatencyOutput{
			P50: helper.FormatDuration(metric.Latency.P50),
			P90: helper.FormatDuration(metric.Latency.P90),
			P95: helper.FormatDuration(metric.Latency.P95),
			P99: helper.FormatDuration(metric.Latency.P99),
		},
		Success:       metric.Success,
		Request:       metric.Request,
		ResponseCodes: metric.ResponseCodes,
		Errors:        metric.Errors,
	}
	jsonByte, err := json.Marshal(metricOutput)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonByte))
	return nil
}

type Latency struct {
	// Mean time.Duration `json:"mean"`
	P50 time.Duration `json:"p50"`
	P90 time.Duration `json:"p90"`
	P95 time.Duration `json:"p95"`
	P99 time.Duration `json:"p99"`

	estimator *tdigest.TDigest
}

type LatencyOutput struct {
	// Mean time.Duration `json:"mean"`
	P50 string `json:"p50"`
	P90 string `json:"p90"`
	P95 string `json:"p95"`
	P99 string `json:"p99"`
}

func (l *Latency) Add(latency time.Duration) {
	l.estimator.Add(float64(latency), 1)
}
