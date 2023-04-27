package client

import (
	"context"

	"github.com/TheLazarusNetwork/erebrus/core"
	"github.com/TheLazarusNetwork/erebrus/model"
	"github.com/TheLazarusNetwork/erebrus/util"
	log "github.com/sirupsen/logrus"
)

// gRPC client service struct
type ClientService struct {
	UnimplementedClientServiceServer
}

// Method to get Client information
func (cs *ClientService) GetClientInformation(ctx context.Context, request *ClientRequest) (*model.Response, error) {
	if ctx.Value("error") == 1 {
		response := core.MakeErrorResponse(500, "Bad Token", nil, nil, nil)
		return response, nil
	}
	id := request.UUID
	log.WithFields(util.StandardFieldsGRPC).Info("Client Information Request ,for:", id)
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

// Method to create client
func (cs *ClientService) CreateClient(ctx context.Context, request *model.Client) (*model.Response, error) {
	if ctx.Value("error") == 1 {
		response := core.MakeErrorResponse(500, "Bad Token", nil, nil, nil)
		return response, nil
	}
	client, err := core.RegisterClient(request)
	log.WithFields(util.StandardFieldsGRPC).Info("Client Creation Request")
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

// Method to update client
func (cs *ClientService) UpdateClient(ctx context.Context, request *UpdateRequest) (*model.Response, error) {
	if ctx.Value("error") == 1 {
		response := core.MakeErrorResponse(500, "Bad Token", nil, nil, nil)
		return response, nil
	}
	id := request.UUID
	log.WithFields(util.StandardFieldsGRPC).Info("Client Update Request ,for:", id)
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

// Method to delete client
func (cs *ClientService) DeleteClient(ctx context.Context, request *ClientRequest) (*model.Response, error) {
	if ctx.Value("error") == 1 {
		response := core.MakeErrorResponse(500, "Bad Token", nil, nil, nil)
		return response, nil
	}
	id := request.UUID
	log.WithFields(util.StandardFieldsGRPC).Info("Delete Client Request ,for:", id)
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

// Method to get all clients
func (cs *ClientService) GetClients(ctx context.Context, request *Empty) (*model.Response, error) {
	if ctx.Value("error") == 1 {
		response := core.MakeErrorResponse(500, "Bad Token", nil, nil, nil)
		return response, nil
	}
	log.WithFields(util.StandardFieldsGRPC).Info("Request For Get All Clients")
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
