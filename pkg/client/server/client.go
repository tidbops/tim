package server

import "github.com/tidbops/tim/pkg/models"

type Client struct{}

func NewServerClient(addr string) (*Client, error) {
	return &Client{}, nil
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

func (c *Client) UpdateTiDBCluster(tc *models.TiDBCluster) error {
	return nil
}
