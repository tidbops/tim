package command

import (
	"io"
	"net/http"
	"os"

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

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
