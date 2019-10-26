package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidbops/tim/pkg/models"
	"net/http"
	"time"
)

type Response struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data []*models.TiDBCluster `json:"data"`
}

func LoadTiDBClusters(c *gin.Context) {
	tc, err := models.LoadTiDBClusters()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("Get failed, %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": tc})
}

func GetTiDBClustersByHost(c *gin.Context) {
	host := c.Query("host")
	if host == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": "host is empty"})
		return
	}
	tc, err := models.GetTiDBClusterByHost(host)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("Get failed, %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": tc})
}

func GetTiDBClustersByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": "name is empty"})
		return
	}
	tc, err := models.GetTiDBClusterByName(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("Get failed, %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": []*models.TiDBCluster{tc}})
}

func CreateTiDBCluster(c *gin.Context) {
	name := c.PostForm("name")
	version := c.PostForm("version")
	path := c.PostForm("path")
	host := c.PostForm("host")
	status := c.PostForm("status")
	dateTime := c.DefaultPostForm("time", time.Now().Format("2006-01-02 15:04:05"))
	if _, err := models.JudgeTiDBStatusType(status); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("TiDBStatus invaild, %v", status)})
		return
	}
	desc := c.PostForm("description")
	t, _ := time.Parse("2006-01-02 15:04:05", dateTime)
	tc := &models.TiDBCluster{
		Name:        name,
		Version:     version,
		Path:        path,
		Host:        host,
		Status:      status,
		Description: desc,
		InitTime:    t,
	}
	if err := models.CreateTiDBCluster(tc); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("store tidb cluster information failed, %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": tc})
}

func SearchTiDBClusters(c *gin.Context) {
	host := c.Query("host")
	name := c.Query("name")
	version := c.Query("version")
	status := c.Query("status")
	path := c.Query("path")

	params := map[string]interface{}{
		"name":    name,
		"version": version,
		"path":    path,
		"host":    host,
		"status":  status,
	}
	tc, err := models.SearchTiDBClusters(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10, "msg": fmt.Sprintf("Search failed, %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": tc})
}
