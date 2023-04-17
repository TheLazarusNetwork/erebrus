package status

import (
	"context"
	"errors"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/util"
	log "github.com/sirupsen/logrus"
)

type StatusService struct {
	UnimplementedStatusServiceServer
}

func (s *StatusService) GetStatus(ctx context.Context, request *Empty) (*model.Status, error) {
	log.WithFields(util.StandardFieldsGRPC).Info("Request For Server Status")
	status, err := core.GetServerStatus()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to get Server Status")
		return nil, errors.New(err.Error())
	}
	return status, nil
}
