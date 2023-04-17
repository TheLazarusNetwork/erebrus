package v1

import (
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/client"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/server"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/status"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {
	//get the instance of server and client services
	ServerService := &server.ServerService{}
	ClientService := &client.ClientService{}
	StatusService := &status.StatusService{}

	//creating a new gRPC server
	grpc_server := grpc.NewServer()

	//Registering server, status and client services
	server.RegisterServerServiceServer(grpc_server, ServerService)
	client.RegisterClientServiceServer(grpc_server, ClientService)
	status.RegisterStatusServiceServer(grpc_server, StatusService)

	return grpc_server
}
