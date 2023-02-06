package paseto

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	gopaseto "aidanwoods.dev/go-paseto"
	log "github.com/sirupsen/logrus"

	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/claims"
	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		err = fmt.Errorf("failed to bind header, %s", err)
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Autherisation header is missing")
		c.Abort()
		return
	}
	token := headers.Authorization
	splitToken := strings.Split(token, "Bearer ")
	pasetoToken := splitToken[1]
	parser := gopaseto.NewParser()
	parser.AddRule(gopaseto.NotExpired())
	publickey := auth.Getpublickey()
	parsedToken, err := parser.ParseV4Public(publickey, pasetoToken, nil)
	jsonvalue := parsedToken.ClaimsJSON()
	ClaimsValue := claims.CustomClaims{}
	json.Unmarshal(jsonvalue, &ClaimsValue)
	if err != nil {
		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bindfailed to scan claims for paseto token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		c.Set("walletAddress", ClaimsValue.WalletAddress)
		c.Next()
	}

}
