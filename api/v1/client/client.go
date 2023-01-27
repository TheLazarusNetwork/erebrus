package client

import (
	"encoding/json"
	"net/http"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		g.POST("/:id", ConfigureClient)
	}
}

// swagger:route POST /client Client createClient
//
// Create client
//
// Create client based on the given client model.
// responses:
//  201: clientSucessResponse
//  400: badRequestResponse
//	401: unauthorizedResponse
//  500: serverErrorResponse

func createClient(c *gin.Context) {
	var data model.Client

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")

		response := core.MakeErrorResponse(400, err.Error(), nil, nil, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client, err := core.CreateClient(&data)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create client")

		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(201, "client created", nil, client, nil)

	c.JSON(http.StatusOK, response)
}

// swagger:route GET /client/{id} Client readClient
//
// # Read client
//
// Return client based on the given uuid.
// responses:
//
//	 200: clientSucessResponse
//	 400: badRequestResponse
//		401: unauthorizedResponse
//	 500: serverErrorResponse
func readClient(c *gin.Context) {
	id := c.Param("id")

	client, err := core.ReadClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read client")

		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "client details", nil, client, nil)

	c.JSON(http.StatusOK, response)
}

// swagger:route PATCH /client/{id} Client updateClient
//
// # Update client
//
// Update client based on the given uuid and client model.
// responses:
//
//	 200: clientSucessResponse
//	 400: badRequestResponse
//		401: unauthorizedResponse
//	 500: serverErrorResponse
func updateClient(c *gin.Context) {
	var data model.Client
	id := c.Param("id")

	if err := c.ShouldBindJSON(&data); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")

		response := core.MakeErrorResponse(400, err.Error(), nil, nil, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client, err := core.UpdateClient(id, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to update client")

		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "client updated", nil, client, nil)

	c.JSON(http.StatusOK, response)
}

// swagger:route DELETE /client/{id} Client deleteClient
//
// # Delete client
//
// Delete client based on the given uuid.
// responses:
//
//	 200: sucessResponse
//	 400: badRequestResponse
//		401: unauthorizedResponse
//	 500: serverErrorResponse
func deleteClient(c *gin.Context) {
	id := c.Param("id")

	err := core.DeleteClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to remove client")

		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "client deleted", nil, nil, nil)

	c.JSON(http.StatusOK, response)
}

// swagger:route GET /client Client readClients
//
// # Read All Clients
//
// Get all clients in the server.
// responses:
//
//	 200: clientsSucessResponse
//	 400: badRequestResponse
//		401: unauthorizedResponse
//	 500: serverErrorResponse
func readClients(c *gin.Context) {
	clients, err := core.ReadClients()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to list clients")

		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := core.MakeSucessResponse(200, "clients details", nil, nil, clients)

	c.JSON(http.StatusOK, response)
}

func ConfigureClient(c *gin.Context) {
	configData, err := core.ReadClientConfig(c.Param("id"))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to read client config")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var responsevalue model.Response
	json.Unmarshal(configData, &responsevalue)
	response := core.MakeSucessResponse(200, "Client Successfully Created", responsevalue.Server, responsevalue.Client, nil)
	c.JSON(http.StatusOK, response)
}
