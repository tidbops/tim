package command

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/bndr/gotabulate"
	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/client/local"
	"github.com/tidbops/tim/pkg/client/server"
	"github.com/tidbops/tim/pkg/models"
)

type Client interface {
	LoadTiDBClusters() ([]*models.TiDBCluster, error)
	GetTiDBClusterByHost(host string) ([]*models.TiDBCluster, error)
	GetTiDBClusterByName(name string) (*models.TiDBCluster, error)
	CreateTiDBCluster(tc *models.TiDBCluster) error
	UpdateTiDBCluster(tc *models.TiDBCluster) error
	SearchTiDBCluster(s map[string]interface{}) ([]*models.TiDBCluster, error)
}

func genClient(cmd *cobra.Command) (Client, error) {
	addr, err := cmd.Flags().GetString("server")
	if err != nil || addr == "" {
		c, err := local.NewLocalClient()
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	c, err := server.NewServerClient(addr)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func initTiDBAnsible(version string, path string) error {
	gitCmd := exec.Command("sh", "-c",
		fmt.Sprintf("git clone -b %s %s %s", version, TiDBAnsibleURL, path))

	stdoutStderr, err := gitCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s, %v", stdoutStderr, err)
	}

	return nil
}

func getHostName() string {
	hostname, _ := os.Hostname()
	return hostname
}

func GetTiDBClustersTableString(tc []*models.TiDBCluster) string {
	var tcArr [][]string
	for _, t := range tc {
		tcArr = append(tcArr, []string{strconv.FormatInt(t.ID, 10), t.Name, t.Version, t.Path, t.Host, t.Status, t.Description, t.InitTime.Format("2006-01-02 15:04:05")})
	}
	t := gotabulate.Create(tcArr)
	t.SetHeaders([]string{"ID", "Name", "Version", "Path", "Host", "Status", "Description", "InitTime"})
	t.SetAlign("right")
	return t.Render("grid")
}
