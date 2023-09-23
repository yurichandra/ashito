package main

import (
	"github.com/spf13/cobra"
	"github.com/yurichandra/ashito/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ashito",
		Short: "Card system load testing application",
		Long:  "ashito is a card system load testing application that run with TCP as underlying connection and using ISO8583 as message format",
	}

	rootCmd.AddCommand(cmd.AttackCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
