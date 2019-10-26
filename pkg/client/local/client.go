package local

import "github.com/tidbops/tim/pkg/models"

type Client struct{}

func NewLocalClient() (*Client, error) {
	c := &Client{}

	if err := models.NewEngine(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) LoadTiDBClusters() ([]*models.TiDBCluster, error) {
	return models.LoadTiDBClusters()
}

func (c *Client) GetTiDBClusterByHost(host string) ([]*models.TiDBCluster, error) {
	return models.GetTiDBClusterByHost(host)
}

func (c *Client) GetTiDBClusterByName(name string) (*models.TiDBCluster, error) {
	return models.GetTiDBClusterByName(name)
}

func (c *Client) CreateTiDBCluster(tc *models.TiDBCluster) error {
	return models.CreateTiDBCluster(tc)
}

func (c *Client) UpdateTiDBCluster(tc *models.TiDBCluster) error {
	return models.UpdateTiDBCluster(tc)
}
