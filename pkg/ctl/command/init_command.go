package command

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	TiDBAnsibleURL = "https://github.com/pingcap/tidb-ansible.git"
)

type InitCommandFlags struct {
	Name        string
	Path        string
	Version     string
	Description string
}

var (
	initCmdFlags = &InitCommandFlags{}
)

func NewInitCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init tidb-ansible files",
		Run:   initCommandFunc,
	}

	initCmd.Flags().StringVar(&initCmdFlags.Name, "name", "", "name specified the name of tidb cluster, required")
	initCmd.Flags().StringVar(&initCmdFlags.Path, "path", "./demo", "path specifies the storage path of the tidb-ansible file, required")
	initCmd.Flags().StringVar(&initCmdFlags.Version, "tidb-version", "master", "version specifies the tidb version to init, required")
	initCmd.Flags().StringVar(&initCmdFlags.Description, "desc", "", "description of the installed tidb cluster")

	return initCmd
}

func initCommandFunc(cmd *cobra.Command, args []string) {
	if initCmdFlags.Name == "" {
		cmd.Println("name flag is required")
		cmd.Println(cmd.UsageString())
		return
	}

	if initCmdFlags.Version == "" {
		cmd.Println("tidb-version flag is required")
		cmd.Println(cmd.UsageString())
		return
	}

	gitCmd := exec.Command("sh", "-c",
		fmt.Sprintf("git clone -b %s %s %s", initCmdFlags.Version, TiDBAnsibleURL, initCmdFlags.Path))

	stdoutStderr, err := gitCmd.CombinedOutput()
	if err != nil {
		cmd.Println(string(stdoutStderr))
		return
	}

	cmd.Printf("Init tidb-ansible files to path %s, version %s", initCmdFlags.Path, initCmdFlags.Version)
}

