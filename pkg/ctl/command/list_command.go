package command

import (
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "tidb-clusters list info",
		Run:   listCommandFunc,
	}
	return listCmd
}

func listCommandFunc(cmd *cobra.Command, args []string) {
	cli, err := genClient(cmd)
	if err != nil {
		cmd.Printf("init client failed, %v", err)
		return
	}
	tc, err := cli.LoadTiDBClusters()
	if err != nil {
		cmd.Printf("load list failed, %v", err)
		return
	}
	if len(tc) == 0 {
		return
	}
	cmd.Println(GetTiDBClustersTableString(tc))
}
