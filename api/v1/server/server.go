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
		g.GET("/status", GetStatus)
	}
}

func readServer(c *gin.Context) {
	server, err := core.ReadServer()
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failure in reading server")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "server details", server, nil, nil)

	c.JSON(http.StatusOK, response)
}

func updateServer(c *gin.Context) {
	var data model.Server

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(util.StandardFields).Error("failed to bind")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	server, err := core.UpdateServer(&data)
	if err != nil {
		log.WithFields(util.StandardFields).Error("failed to update server")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "server updated", server, nil, nil)

	c.JSON(http.StatusOK, response)
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

func GetStatus(c *gin.Context) {
	status_data, err := core.GetServerStatus()
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failed to get server status")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, status_data)
}
