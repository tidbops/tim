package ctl

import (
	"gopkg.in/mikefarah/yaml.v2"
	"os"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/ctl/command"
)

var (
	url string
)

func init() {
	cobra.EnablePrefixMatching = true
	yaml.DefaultMapType = reflect.TypeOf(yaml.MapSlice{})
}

func Start(args []string) {
	rootCmd := &cobra.Command{
		Use:        "tim",
		Short:      "TiM is a tool for managing multiple tidb clusters",
		Long:       "A tool to manage multi tidb-ansible and help to upgrade tidb version",
		SuggestFor: []string{"tim-ctl"},
	}

	rootCmd.Flags().StringVarP(&url, "server", "u", "", "tim-server address")

	rootCmd.AddCommand(
		command.NewInitCommand(),
		command.NewUpgradeCommand(),
		command.NewListCommand(),
		command.NewSearchCommand(),
		command.NewEnvCommand(),
	)

	rootCmd.SetArgs(args)
	rootCmd.SilenceErrors = true
	rootCmd.ParseFlags(args)
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println(err)
	}
}
