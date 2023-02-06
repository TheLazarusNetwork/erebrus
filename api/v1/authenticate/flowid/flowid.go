package flowid

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
	RelatedRoleId string
}
type Db struct {
	WalletAddress string
	Timestamp     time.Time
}

var data map[string]Db

// ApplyRoutes applies router to gin Router
// func ApplyRoutes(r *gin.RouterGroup) {
// 	g := r.Group("/flowid")
// 	{
// 		g.GET("", GetFlowId)
// 	}
// }

func GetFlowId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" {
		log.WithFields(log.Fields{
			"err": "empty Wallet Address",
		}).Error("failed to create client")

		response := core.MakeErrorResponse(500, "Empty Wallet Address", nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
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
	flowId, err := GenerateFlowId(walletAddress, "")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create FlowId")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userAuthEULA := os.Getenv("AUTH_EULA")
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	c.JSON(200, payload)
}

func GenerateFlowId(walletAddress string, relatedRoleId string) (string, error) {

	flowId := uuid.NewString()
	var dbdata Db
	dbdata.WalletAddress = walletAddress
	dbdata.Timestamp = time.Now()
	data = map[string]Db{
		flowId: dbdata,
	}
	return flowId, nil
}
