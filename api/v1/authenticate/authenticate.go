package authenticate

import (
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/erebrus/api/v1/authenticate/challengeid"
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
		g.GET("", challengeid.GetChallengeId)
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
	message := userAuthEULA + req.ChallengeId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.ChallengeId, message)

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
		delete(challengeid.Data, req.ChallengeId)
		payload := AuthenticatePayload{
			Status:  200,
			Success: true,
			Message: "Successfully Authenticated",
			Token:   pasetoToken,
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
	payload.Success = false
	payload.Status = 401
	payload.Message = errvalue
	return payload
}
