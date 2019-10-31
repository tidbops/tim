package command

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/models"
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

	if err := initTiDBAnsible(initCmdFlags.Version, initCmdFlags.Path); err != nil {
		cmd.Println(err)
		return
	}

	tc := &models.TiDBCluster{
		Name:        initCmdFlags.Name,
		Version:     initCmdFlags.Version,
		Path:        initCmdFlags.Path,
		Description: initCmdFlags.Description,
		InitTime:    time.Now(),
		Host:        getHostName(),
		Status:      string(models.TiDBInited),
	}
	cli, err := genClient(cmd)
	if err != nil {
		cmd.Printf("init client failed, %v\n", err)
		return
	}

	if err := cli.CreateTiDBCluster(tc); err != nil {
		cmd.Printf("store tidb cluster information failed, %v\n", err)
		return
	}
	cmd.Printf("Success! tidb-ansible files saved %s, version %s\n", initCmdFlags.Path, initCmdFlags.Version)
}
