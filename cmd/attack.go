package cmd

import "github.com/spf13/cobra"

var AttackCmd = &cobra.Command{
	Use:   "attack",
	Short: "attack the target",
	Long:  "attack will attacks the target flag with input from the file",
}
