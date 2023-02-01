package flowid

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/dbconfig"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Name          string   `json:"name,omitempty"`
	WalletAddress string   `gorm:"primary_key" json:"walletAddress"`
	FlowIds       []FlowId `gorm:"foreignkey:WalletAddress" json:"-"`
}
type FlowId struct {
	WalletAddress string
	FlowId        string `gorm:"primary_key"`
	RelatedRoleId string
}

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.GET("", GetFlowId)
	}
}

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
		c.JSON(http.StatusInternalServerError, response)
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
	db := dbconfig.GetDb()
	flowId := uuid.NewString()
	var update bool
	update = true

	findResult := db.Model(&User{}).Find(&User{}, &User{WalletAddress: walletAddress})

	if err := findResult.Error; err != nil {
		err = fmt.Errorf("while finding user error occured, %s", err)
		logrus.Error(err)
		return "", err
	}

	rowsAffected := findResult.RowsAffected
	if rowsAffected == 0 {
		update = false
	}
	if update {
		// User exist so update
		association := db.Model(&User{
			WalletAddress: walletAddress,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		err := association.Append(&FlowId{WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId})
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create

		newUser := &User{
			WalletAddress: walletAddress,
			FlowIds: []FlowId{{
				WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
