package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngaut/log"
	"github.com/tidbops/tim/pkg/models"
	"github.com/tidbops/tim/pkg/server"
)

func main() {
	g := gin.Default()

	server.Router(g)

	if err := models.NewEngine(); err != nil {
		log.Fatal(err)
	}
	// Listen and serve on 0.0.0.0:8080
	g.Run(":8080")
}
