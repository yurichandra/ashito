package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yurichandra/ashito/internal"
)

var AttackCmd = &cobra.Command{
	Use:   "attack",
	Short: "attack the target",
	Long:  "attack will attacks the target flag with input from the file",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	inputFlag := cmd.Flag("input")
	filePath := inputFlag.Value.String()
	if filePath == "" {
		log.Println("flag `input` is required")
		return
	}

	err := internal.Attack(filePath)
	if err != nil {
		fmt.Println(err)
		// do nothing as of now
	}
}
