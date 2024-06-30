package internal

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/yurichandra/ashito/helper"
	"github.com/yurichandra/ashito/message"
)

type transaction struct {
	cardNumber string
	amount     string
	currency   string
}

type Flag struct {
	FilePath string
	Target   string
	// In second
	Duration int
}

type connectionPools struct {
	connections []net.Conn
}

type Metrics struct {
	// Success rate of requests
	Success float64 `json:"success"`
	// Total number of requests made
	Request int `json:"request"`
	// Map of response codes
	ResponseCodes map[string]int `json:"response_codes"`
	// Collection of errors on attacks
	Errors []string `json:"errors"`
}

func Attack(flag Flag) error {
	file, err := os.Open(flag.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	transactions := make([]transaction, len(records)-1)
	for index, data := range records {
		if index == 0 {
			continue
		}

		transactions[index-1] = transaction{
			cardNumber: data[0],
			amount:     data[1],
			currency:   data[2],
		}
	}

	requestBytes, err := prepareMessage(transactions)
	if err != nil {
		return err
	}

	connectionPools, err := prepareConnectionPools(len(transactions), flag.Target)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	// Create channels for passing data from goroutines into main goroutine
	errChan := make(chan error)
	responseChan := make(chan []byte)
	doneChan := make(chan bool)

	wg := &sync.WaitGroup{}
	for index, requestByte := range requestBytes {
		wg.Add(1)
		connection := connectionPools.connections[index]

		go func(ctx context.Context, requestByte []byte, conn net.Conn) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					err := helper.WriteOnTopTCP(conn, requestByte)
					if err != nil {
						errChan <- err
					}

					responseByte, err := helper.ReadOnTopTCP(bufio.NewReader(conn))
					if err != nil {
						errChan <- err
					}

					if responseByte != nil {
						responseChan <- responseByte
					}
				}
			}

		}(ctx, requestByte, connection)
	}

	go func() {
		wg.Wait()
		doneChan <- true
	}()

	metrics := &Metrics{
		Success:       0,
		Request:       0,
		ResponseCodes: make(map[string]int),
		Errors:        make([]string, 0),
	}

	for {
		select {
		case err := <-errChan:
			metrics.Errors = append(metrics.Errors, err.Error())
			metrics.Request++
		case response := <-responseChan:
			metrics.Request++
			resp, err := message.UnpackMessage(message.CbsSpec, response)
			if err != nil {
				metrics.Errors = append(metrics.Errors, err.Error())
				continue
			}

			responseCode, err := resp.GetField(39).String()
			if err != nil {
				metrics.Errors = append(metrics.Errors, err.Error())
				continue
			}

			metrics.ResponseCodes[responseCode]++

			if responseCode == "000" {
				metrics.Success++
			}
		case <-doneChan:
			metrics.Success = (metrics.Success / float64(metrics.Request)) * 100

			// By default, printing result into a json
			jsonByte, err := json.Marshal(metrics)
			if err != nil {
				return err
			}

			fmt.Println(string(jsonByte))
			return nil
		}
	}
}

func prepareConnectionPools(numOfPools int, target string) (connectionPools, error) {
	counter := 0
	pools := make([]net.Conn, numOfPools)

	for counter < numOfPools {
		conn, err := net.Dial("tcp", target)
		if err != nil {
			return connectionPools{}, err
		}

		pools[counter] = conn
		counter++
	}

	return connectionPools{
		connections: pools,
	}, nil
}

func prepareMessage(transactions []transaction) ([][]byte, error) {
	requestBytes := make([][]byte, len(transactions))
	template, err := helper.UnpackBytesFromString([]byte(message.CBSPurchaseRequestTemplate))
	if err != nil {
		return nil, err
	}

	for index, transaction := range transactions {
		templateMessage, err := message.UnpackMessage(message.CbsSpec, template)
		if err != nil {
			return nil, err
		}

		// For now only set card number, amount, and currency, expand the capability if more
		// fields need to be set.
		templateMessage.Field(2, transaction.cardNumber)
		templateMessage.Field(4, transaction.amount)
		templateMessage.Field(37, helper.StringNumber(message.RrnLength))
		templateMessage.Field(49, transaction.currency)

		finalBytes, err := templateMessage.Pack()
		if err != nil {
			return nil, err
		}

		requestBytes[index] = finalBytes
	}

	return requestBytes, nil
}
