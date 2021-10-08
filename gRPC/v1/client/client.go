package client

import (
	"context"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	log "github.com/sirupsen/logrus"
)

//gRPC client service struct
type ClientService struct {
	UnimplementedClientServiceServer
}

//Method to get Client information
func (cs *ClientService) GetClientInformation(ctx context.Context, request *ClientRequest) (*model.Response, error) {
	id := request.UUID

	client, err := core.ReadClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Client Information Fetched", nil, client, nil)
	return response, nil
}

//Method to get client config data
func (cs *ClientService) GetClientConfiguration(ctx context.Context, request *ClientRequest) (*Config, error) {
	id := request.UUID
	configData, err := core.ReadClientConfig(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client config")
		return nil, err
	}

	return &Config{Config: configData}, nil
}

// Method to email client the configuration file
func (cs *ClientService) EmailClientConfiguration(ctx context.Context, request *ClientRequest) (*model.Response, error) {
	id := request.UUID

	err := core.EmailClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Client Configuration Emailed", nil, nil, nil)
	return response, nil
}

//Method to create client
func (cs *ClientService) CreateClient(ctx context.Context, request *model.Client) (*model.Response, error) {
	client, err := core.CreateClient(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(201, "Client Created", nil, client, nil)
	return response, nil
}

//Method to update client
func (cs *ClientService) UpdateClient(ctx context.Context, request *UpdateRequest) (*model.Response, error) {
	id := request.UUID
	client, err := core.UpdateClient(id, request.Client)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Client Updated", nil, client, nil)
	return response, nil
}

//Method to delete client
func (cs *ClientService) DeleteClient(ctx context.Context, request *ClientRequest) (*model.Response, error) {
	id := request.UUID
	err := core.DeleteClient(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Client Deleted", nil, nil, nil)
	return response, nil
}

//Method to get all clients
func (cs *ClientService) GetClients(ctx context.Context, request *Empty) (*model.Response, error) {
	clients, err := core.ReadClients()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("unable to read client")
		response := core.MakeErrorResponse(500, err.Error(), nil, nil, nil)
		return response, err
	}

	response := core.MakeSucessResponse(200, "Client Information Fetched", nil, nil, clients)
	return response, nil
}
