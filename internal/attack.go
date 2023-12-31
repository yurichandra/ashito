package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type transaction struct {
	cardNumber string
	amount     string
	currency   string
}

func Attack(filePath string) error {
	file, err := os.Open(filePath)
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

	fmt.Println(transactions)
	return nil
}
