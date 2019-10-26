package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tidbops/tim/pkg/server/api"
	"github.com/tidbops/tim/pkg/server/dashboard/templates"
	"github.com/tidbops/tim/pkg/server/dashboard/web"
	"html/template"
	"io/ioutil"
	"strings"
)

func Router(r *gin.Engine) {
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("api/status", api.Status)
	r.GET("api/loadtidbclusters", api.LoadTiDBClusters)
	r.POST("api/createtidbcluster", api.CreateTiDBCluster)
	r.GET("api/gettidbclustersbyname", api.GetTiDBClustersByName)
	r.GET("api/gettidbclustersbyhost", api.GetTiDBClustersByHost)
	r.GET("api/searchtidbclusters", api.SearchTiDBClusters)

	r.GET("dashboard/index", web.Index)
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range templates.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
