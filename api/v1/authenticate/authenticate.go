package authenticate

import (
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/erebrus/api/v1/authenticate/flowid"
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

		errResponse := ErrAuthenticate(err.Error())
		c.JSON(http.StatusForbidden, errResponse)
		return
	}
	userAuthEULA := os.Getenv("AUTH_EULA")
	message := userAuthEULA + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("FlowId Not Found")
		errResponse := ErrAuthenticate(err.Error())
		c.JSON(http.StatusNotFound, errResponse)
		return
	}

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to CheckSignature")
		errResponse := ErrAuthenticate(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		pasetoToken, err := auth.GenerateTokenPaseto(customClaims)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to generate token")
			errResponse := ErrAuthenticate(err.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		delete(flowid.Data, req.FlowId)
		payload := AuthenticatePayload{
			StatusDesc: "Successfully Authenticated",
			Token:      pasetoToken,
		}
		c.JSON(http.StatusAccepted, payload)
	} else {
		errResponse := ErrAuthenticate("Forbidden")
		c.JSON(http.StatusForbidden, errResponse)
		return
	}
}

func ErrAuthenticate(errvalue string) AuthenticatePayload {
	var payload AuthenticatePayload
	payload.StatusDesc = errvalue
	payload.Token = "nil"
	return payload
}
