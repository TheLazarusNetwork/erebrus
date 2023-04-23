package paseto

import (
	context "context"
	"encoding/json"
	"errors"
	"fmt"

	gopaseto "aidanwoods.dev/go-paseto"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/claims"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func PASETO(ctx context.Context) (context.Context, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md["authorization"][0]
	if token == "" {
		log.WithFields(log.Fields{
			"err": "Authorization header is missing",
		}).Error("Authorization header is missing")
		return nil, errors.New("authorization header is missing")
	}
	parser := gopaseto.NewParser()
	parser.AddRule(gopaseto.NotExpired())
	publickey := auth.Getpublickey()
	parsedToken, err := parser.ParseV4Public(publickey, token, nil)
	if err != nil {
		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bindfailed to scan claims for paseto token")
		return nil, errors.New(err.Error())
	}
	jsonvalue := parsedToken.ClaimsJSON()
	ClaimsValue := claims.CustomClaims{}
	json.Unmarshal(jsonvalue, &ClaimsValue)
	new_ctx := context.WithValue(ctx, "walletAddress", ClaimsValue.WalletAddress)

	return new_ctx, nil
}
