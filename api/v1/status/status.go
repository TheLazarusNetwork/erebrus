package status

import (
	"net/http"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	r.GET("/status", GetStatus)
}

// swagger:route GET /server/status Server statusServer
//
// Get Server status
//
// Retrieves the server  status details.
// responses:
//  200: serverStatusResponse
//  400: badRequestResponse
//	401: unauthorizedResponse
//  500: serverErrorResponse
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
