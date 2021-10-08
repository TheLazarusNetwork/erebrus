package server

import (
	"context"
	"errors"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/util"
	log "github.com/sirupsen/logrus"
)

type ServerService struct {
	UnimplementedServerServiceServer
}

//Method to get server information
func (ss *ServerService) GetServerInformation(ctx context.Context, request *Empty) (*model.Response, error) {
	server, err := core.ReadServer()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to get server info")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Server Information Fetched", server, nil, nil)
	return response, nil
}

//method to get server configuration
func (ss *ServerService) GetServerConfiguraion(ctx context.Context, request *Empty) (*Config, error) {
	configData, err := core.ReadWgConfigFile()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read server configuration")

		return nil, errors.New(err.Error())
	}

	return &Config{Config: configData}, nil
}

//Method to get server status
func (ss *ServerService) GetStatus(ctx context.Context, request *Empty) (*model.Status, error) {
	status, err := core.GetServerStatus()
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failed to get server status")
		return nil, errors.New(err.Error())
	}
	return status, nil
}

//Method to update server
func (ss *ServerService) UpdateServer(ctx context.Context, request *model.Server) (*model.Response, error) {
	server, err := core.UpdateServer(request)
	if err != nil {
		log.WithFields(util.StandardFields).Error("Failed to update server")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Server Updated", server, nil, nil)
	return response, nil
}
