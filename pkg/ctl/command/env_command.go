package command

import (
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/utils"
)

func NewEnvCommand() *cobra.Command {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "init environment for tidb-ansible",
		Run:   envCommandFunc,
	}

	return envCmd
}

const (
	initScriptURL = "https://raw.githubusercontent.com/tidbops/tim/master/scripts/init_env.sh"
	initScitpFile = "/tmp/init_env.sh"
)

func envCommandFunc(cmd *cobra.Command, args []string) {
	if err := utils.DownloadFile(initScriptURL, initScitpFile); err != nil {
		cmd.Println(err)
		return
	}

	shCmd := exec.Command("sh", initScitpFile)
	stdoutStderr, err := shCmd.CombinedOutput()
	if err != nil {
		cmd.Println(string(stdoutStderr))
		return
	}

	cmd.Println("Success!")
}
