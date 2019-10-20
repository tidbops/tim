package ctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/ctl/command"
)

var rootCmd = &cobra.Command{
	Use:        "tim",
	Short:      "TiM is a tool for managing multiple tidb clusters",
	Long:       "A tool to manage multi tidb-ansible and help to upgrade tidb version",
	SuggestFor: []string{"tim-ctl"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("shell command")
		}
	},
}

func init() {
	rootCmd.AddCommand(
		command.NewInitCommand(),
	)
}

// Execute execs Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		command.ExitWithError(command.ExitError, err)
	}
}
