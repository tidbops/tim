package command

import (
	"github.com/spf13/cobra"
)

type InitCommandFlags struct {
	Name        string
	Path        string
	Version     string
	Description string
}

var (
	initCmdFlags = InitCommandFlags{}
)

func NewInitCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init tidb cluster",
		Run:   initCommandFunc,
	}

	initCmd.Flags().StringVar(&initCmdFlags.Name, "name", "", "name specified the name of tidb cluster, required")
	initCmd.Flags().StringVar(&initCmdFlags.Path, "path", ".", "path specifies the storage path of the tidb-ansible file, required")
	initCmd.Flags().StringVar(&initCmdFlags.Version, "version", "latest", "version specifies the tidb version to init, required")
	initCmd.Flags().StringVar(&initCmdFlags.Description, "desc", "", "description of the installed tidb cluster")

	if err := initCmd.MarkFlagRequired("name"); err != nil {
		ExitWithError(ExitError, err)
	}

	return initCmd
}

func initCommandFunc(cmd *cobra.Command, args []string) {

}
