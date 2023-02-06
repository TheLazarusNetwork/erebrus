package authenticate

import (
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/erebrus/api/v1/authenticate/flowid"
	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/claims"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/cryptosign"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.GET("", flowid.GetFlowId)
		g.POST("", authenticate)

	}
}

func authenticate(c *gin.Context) {

	var req AuthenticateRequest
	err := c.BindJSON(&req)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Invalid request payload")

		response := core.MakeErrorResponse(400, err.Error(), nil, nil, nil)
		c.JSON(http.StatusForbidden, response)
		return
	}
	//Get flowid type
	// var flowIdData flowid.FlowId
	// err = db.Model(&flowid.FlowId{}).Where("flow_id = ?", req.FlowId).First(&flowIdData).Error
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"err": err,
	// 	}).Error("Flow Id Not found")

	// 	response := core.MakeErrorResponse(404, err.Error(), nil, nil, nil)
	// 	c.JSON(http.StatusNotFound, response)

	// 	return
	// }
	// localData, exists := flowid.Data[req.FlowId]
	// if !exists {
	// 	log.WithFields(log.Fields{
	// 		"err": "Flow Id Not Found",
	// 	}).Error("Flow Id Not found")

	// 	response := core.MakeErrorResponse(404, "Flow Id Not Found", nil, nil, nil)
	// 	c.JSON(http.StatusNotFound, response)
	// }

	userAuthEULA := os.Getenv("AUTH_EULA")
	message := userAuthEULA + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {

		c.JSON(http.StatusNotFound, "nil")
		return
	}

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to CheckSignature")

		c.JSON(http.StatusInternalServerError, "nil")
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		pasetoToken, err := auth.GenerateTokenPaseto(customClaims)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to generate token")

			c.JSON(http.StatusInternalServerError, "nil")
			return
		}
		delete(flowid.Data, req.FlowId)
		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		c.JSON(http.StatusAccepted, payload)
	} else {
		c.JSON(http.StatusForbidden, "nil")
		return
	}
}
