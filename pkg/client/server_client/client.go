package server_client

import "github.com/tidbops/tim/pkg/models"

type Client struct{}

func NewClient() *Client {
}

func (c *Client) LoadTiDBClusters() ([]*models.TiDBCluster, error) {
	return nil, nil
}

func (c *Client) GetTiDBClusterByHost(host string) ([]*models.TiDBCluster, error) {
	return nil, nil
}

func (c *Client) GetTiDBClusterByName(name string) (*models.TiDBCluster, error) {
	return nil, nil
}

func (c *Client) CreateTiDBCluster(tc *models.TiDBCluster) error {
	return nil
}
