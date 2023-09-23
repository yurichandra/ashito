package main

import (
	"github.com/spf13/cobra"
	"github.com/yurichandra/ashito/cmd"
)

var Input string

func main() {
	rootCmd := &cobra.Command{
		Use:   "ashito",
		Short: "Card system load testing application",
		Long:  "ashito is a card system load testing application that run with TCP as underlying connection and using ISO8583 as message format",
	}

	attackCommand := cmd.AttackCmd
	attackCommand.Flags().StringVarP(&Input, "input", "i", "", "Input file that have cards and transaction data, file format should be .csv")

	rootCmd.AddCommand(attackCommand)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
