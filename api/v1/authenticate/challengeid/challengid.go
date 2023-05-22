package challengeid

import (
	"net/http"
	"os"
	"time"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/util"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

//	type User struct {
//		Name          string   `json:"name,omitempty"`
//		WalletAddress string   `gorm:"primary_key" json:"walletAddress"`
//		FlowIds       []FlowId `gorm:"foreignkey:WalletAddress" json:"-"`
//	}
type FlowId struct {
	WalletAddress string
	FlowId        string `gorm:"primary_key"`
}
type Db struct {
	WalletAddress string
	Timestamp     time.Time
}

var Data map[string]Db

// ApplyRoutes applies router to gin Router
// func ApplyRoutes(r *gin.RouterGroup) {
// 	g := r.Group("/flowid")
// 	{
// 		g.GET("", GetFlowId)
// 	}
// }

func GetChallengeId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" {
		log.WithFields(log.Fields{
			"err": "empty Wallet Address",
		}).Error("failed to create client")

		response := core.MakeErrorResponse(403, "Empty Wallet Address", nil, nil, nil)
		c.JSON(http.StatusForbidden, response)
		return
	}
	_, err := hexutil.Decode(walletAddress)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Wallet address (walletAddress) is not valid")

		response := core.MakeErrorResponse(400, err.Error(), nil, nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if !util.RegexpWalletEth.MatchString(walletAddress) {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Wallet address (walletAddress) is not valid")
		response := core.MakeErrorResponse(400, err.Error(), nil, nil, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	challengeId, err := GenerateChallengeId(walletAddress)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create FlowId")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	isAuthorized := walletAddress == os.Getenv("MASTERNODE_WALLET")
	userAuthEULA := os.Getenv("AUTH_EULA")
	payload := GetChallengeIdPayload{
		ChallengeId:  challengeId,
		Eula:         userAuthEULA,
		IsAuthorized: isAuthorized,
	}
	c.JSON(200, payload)
}

func GenerateChallengeId(walletAddress string) (string, error) {

	challengeId := uuid.NewString()
	var dbdata Db
	dbdata.WalletAddress = walletAddress
	dbdata.Timestamp = time.Now()
	Data = map[string]Db{
		challengeId: dbdata,
	}
	return challengeId, nil
}
