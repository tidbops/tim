package server

import (
	"encoding/json"
	"fmt"
	"github.com/tidbops/tim/pkg/models"
	"github.com/tidbops/tim/pkg/server/api"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct{}

var (
	address string
)

func NewServerClient(addr string) (*Client, error) {
	address = addr
	return &Client{}, nil
}

func (c *Client) LoadTiDBClusters() ([]*models.TiDBCluster, error) {
	resp, err := getRpcCall("/api/loadtidbclusters", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func (c *Client) GetTiDBClusterByHost(host string) ([]*models.TiDBCluster, error) {
	params := map[string]interface{}{
		"host": host,
	}
	resp, err := getRpcCall("/api/gettidbclustersbyhost", params)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func (c *Client) GetTiDBClusterByName(name string) (*models.TiDBCluster, error) {
	params := map[string]interface{}{
		"name": name,
	}
	resp, err := getRpcCall("/api/gettidbclustersbyname", params)
	if err != nil {
		return nil, err
	}
	return resp.Data[0], err
}

func (c *Client) CreateTiDBCluster(tc *models.TiDBCluster) error {
	params := map[string]interface{}{
		"name":        tc.Name,
		"version":     tc.Version,
		"path":        tc.Path,
		"host":        tc.Host,
		"status":      tc.Status,
		"description": tc.Description,
		//"initTime":    tc.InitTime,
	}
	_, err := postRpcCall("/api/createtidbcluster", params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTiDBCluster(tc *models.TiDBCluster) error {
	return nil
}

func (c *Client) SearchTiDBCluster(s map[string]interface{}) ([]*models.TiDBCluster, error) {
	resp, err := getRpcCall("/api/searchtidbclusters", s)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func getRpcCall(apiMethod string, params map[string]interface{}) (*api.Response, error) {
	p := ""
	for k, v := range params {
		if v != "" {
			p += k + "=" + v.(string) + "&"
		}
	}
	if p != "" {
		p = "?" + strings.TrimSuffix(p, "&")
	}
	resp, err := http.Get("http://" + address + apiMethod + p)
	if err != nil {
		return nil, fmt.Errorf("get call failed, %v", err)
	}
	return parseResponse(resp)
}

func postRpcCall(apiMethod string, params map[string]interface{}) (*api.Response, error) {
	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Set(k, v.(string))
		}
	}
	resp, err := http.PostForm("http://"+address+apiMethod, values)
	if err != nil {
		return nil, fmt.Errorf("call failed, %v", err)
	}
	return parseResponse(resp)
}

func parseResponse(resp *http.Response) (*api.Response, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read failed, %v", err)
	}

	respBody := &api.Response{}
	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, fmt.Errorf("jsonUnmarshal failed, %v", err)
	}

	if respBody.Code != 0 {
		return nil, fmt.Errorf("code not 0, %s", respBody.Msg)
	}
	return respBody, nil
}
