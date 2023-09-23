package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yurichandra/ashito/internal"
)

var AttackCmd = &cobra.Command{
	Use:   "attack",
	Short: "attack the target",
	Long:  "attack will attacks the target flag with input from the file",
	Run:   run,
}

var totalDefaultWorker = 10

func run(cmd *cobra.Command, args []string) {
	inputFlag := cmd.Flag("input")
	filePath := inputFlag.Value.String()
	if filePath == "" {
		log.Println("flag `input` is required")
		return
	}

	targetFlag := cmd.Flag("target")
	target := targetFlag.Value.String()
	if target == "" {
		log.Println("flag `target` is required")
		return
	}

	flag := internal.Flag{
		FilePath:  filePath,
		Target:    target,
		MaxWorker: totalDefaultWorker,
	}

	maxWorkerFlag := cmd.Flag("worker")
	maxWorker := maxWorkerFlag.Value.String()
	if maxWorker != "" {
		parsedMaxWorker, err := strconv.Atoi(maxWorker)
		if err != nil {
			log.Println(err)
			return
		}

		flag.MaxWorker = parsedMaxWorker
	}

	err := internal.Attack(flag)
	if err != nil {
		fmt.Println(err)
		// do nothing as of now
	}
}
