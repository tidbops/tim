package command

import (
	"github.com/spf13/cobra"
)

type SearchCommandFlags struct {
	Name    string
	Path    string
	Version string
	Status  string
	Host    string
}

var (
	searchCmdFlags = &SearchCommandFlags{}
)

func NewSearchCommand() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "tidb-clusters search info",
		Run:   searchCommandFunc,
	}
	searchCmd.Flags().StringVar(&searchCmdFlags.Name, "n", "", "the name of tidb cluster")
	searchCmd.Flags().StringVar(&searchCmdFlags.Path, "p", "", "the storage path of the tidb-ansible file")
	searchCmd.Flags().StringVar(&searchCmdFlags.Version, "v", "", "the tidb version")
	searchCmd.Flags().StringVar(&searchCmdFlags.Status, "s", "", "the tidb-cluster status")
	searchCmd.Flags().StringVar(&searchCmdFlags.Host, "h", "", "the tidb-cluster host")
	return searchCmd
}

func searchCommandFunc(cmd *cobra.Command, args []string) {
	flags := map[string]interface{}{
		"name":    searchCmdFlags.Name,
		"path":    searchCmdFlags.Path,
		"version": searchCmdFlags.Version,
		"status":  searchCmdFlags.Status,
		"host":    searchCmdFlags.Host,
	}
	cli, err := genClient(cmd)
	if err != nil {
		cmd.Printf("init client failed, %v", err)
		return
	}
	tc, err := cli.SearchTiDBCluster(flags)
	if err != nil {
		cmd.Printf("search failed, %v", err)
		return
	}
	if len(tc) == 0 {
		return
	}
	cmd.Println(GetTiDBClustersTableString(tc))
}
