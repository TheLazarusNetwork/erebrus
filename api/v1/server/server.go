package server

import (
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/server")
	{
		g.GET("", readServer)
		g.PATCH("", updateServer)
		g.GET("/config", configServer)
		g.GET("/version", versionStr)
	}
}

func readServer(c *gin.Context) {
	server, err := core.ReadServer()
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failure in reading server")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, server)
}

func updateServer(c *gin.Context) {
	var data model.Server

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(util.StandardFields).Error("failed to bind")
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	server, err := core.UpdateServer(&data)
	if err != nil {
		log.WithFields(util.StandardFields).Error("failed to update server")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, server)
}

func configServer(c *gin.Context) {
	configData, err := core.ReadWgConfigFile()
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failed to read wireguard config file")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// return config as txt file
	c.Header("Content-Disposition", "attachment; filename="+os.Getenv("WG_INTERFACE_NAME")+"")
	c.Data(http.StatusOK, "application/config", configData)
}

func versionStr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": util.Version,
	})
}
