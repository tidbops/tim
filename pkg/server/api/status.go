package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Msg": "OK"})
}
