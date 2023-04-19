package paseto

import (
	context "context"
	"encoding/json"
	"errors"
	"fmt"

	gopaseto "aidanwoods.dev/go-paseto"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/auth"
	"github.com/TheLazarusNetwork/erebrus/util/pkg/claims"
	authMiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func PASETO(ctx context.Context) (context.Context, error) {
	token, err := authMiddleware.AuthFromMD(ctx, "bearer")
	if token == "" {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Autherisation header is missing")
		return nil, errors.New(err.Error())
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
	} else {
		jsonvalue := parsedToken.ClaimsJSON()
		ClaimsValue := claims.CustomClaims{}
		json.Unmarshal(jsonvalue, &ClaimsValue)
		header := metadata.Pairs("header-key", ClaimsValue.WalletAddress)
		grpc.SendHeader(ctx, header)
	}
	return context.Context(ctx), nil
}
