package main

import (
	"github.com/spf13/cobra"
	"github.com/yurichandra/ashito/cmd"
)

var (
	Input  string
	Worker string
	Target string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ashito",
		Short: "Card system load testing application",
		Long:  "ashito is a card system load testing application that run with TCP as underlying connection and using ISO8583 as message format",
	}

	attackCommand := cmd.AttackCmd
	attackCommand.Flags().StringVarP(&Input, "input", "i", "", "Input file that have cards and transaction data, file format should be .csv")
	attackCommand.Flags().StringVarP(&Worker, "worker", "w", "10", "Number of worker")
	attackCommand.Flags().StringVarP(&Target, "target", "", "", "Target destination host")

	rootCmd.AddCommand(attackCommand)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
