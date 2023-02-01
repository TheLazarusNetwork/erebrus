package claims

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TheLazarusNetwork/erebrus/api/v1/flowid"
	"github.com/TheLazarusNetwork/erebrus/dbconfig"
	log "github.com/sirupsen/logrus"
)

type CustomClaims struct {
	WalletAddress string    `json:"walletAddress"`
	SignedBy      string    `json:"signedBy"`
	Expiration    time.Time `json:"expiryTime"`
}

func (c CustomClaims) Valid() error {
	db := dbconfig.GetDb()
	err := db.Model(&flowid.User{}).Where("wallet_address = ?", c.WalletAddress).First(&flowid.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func New(walletAddress string) CustomClaims {
	pasetoExpirationInHours, ok := os.LookupEnv("PASETO_EXPIRATION_IN_HOURS")
	pasetoExpirationInHoursInt := time.Duration(24)
	fmt.Println("ok value walletaddress", ok)
	if ok {
		res, err := strconv.Atoi(pasetoExpirationInHours)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to bind")
		} else {
			pasetoExpirationInHoursInt = time.Duration(res)
		}
	}
	pasetoExpirationHours := pasetoExpirationInHoursInt * time.Hour
	expiration := time.Now().Add(pasetoExpirationHours)
	signedBy := os.Getenv("SIGNED_BY")
	return CustomClaims{
		walletAddress,
		signedBy,
		expiration,
	}
}
