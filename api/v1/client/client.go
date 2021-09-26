package client

import (
	"net/http"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/client")
	{

		g.POST("", createClient)
		g.GET("/:id", readClient)
		g.PATCH("/:id", updateClient)
		g.DELETE("/:id", deleteClient)
		g.GET("", readClients)
		g.GET("/:id/config", configClient)
		g.GET("/:id/email", emailClient)
	}
}

func createClient(c *gin.Context) {
	var data model.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	client, err := core.CreateClient(&data)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create client")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, client)
}

func readClient(c *gin.Context) {
	id := c.Param("id")

	client, err := core.ReadClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read client")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, client)
}

func updateClient(c *gin.Context) {
	var data model.Client
	id := c.Param("id")

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	client, err := core.UpdateClient(id, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to update client")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, client)
}

func deleteClient(c *gin.Context) {
	id := c.Param("id")

	err := core.DeleteClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to remove client")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func readClients(c *gin.Context) {
	clients, err := core.ReadClients()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to list clients")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, clients)
}

func configClient(c *gin.Context) {
	configData, err := core.ReadClientConfig(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read client config")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	formatQr := c.DefaultQuery("qrcode", "false")
	if formatQr == "false" {
		// return config as txt file
		c.Header("Content-Disposition", "attachment; filename="+c.Param("id")+".conf")
		c.Data(http.StatusOK, "application/config", configData)
		return
	}
	// return config as png qrcode
	png, err := qrcode.Encode(string(configData), qrcode.Medium, 250)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create qrcode")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, "image/png", png)
	return
}

func emailClient(c *gin.Context) {
	id := c.Param("id")

	err := core.EmailClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to send email to client")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
