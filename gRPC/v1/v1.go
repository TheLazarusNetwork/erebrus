package v1

import (
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/client"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/server"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {
	//get the instance of server and client services
	ServerService := &server.ServerService{}
	ClientService := &client.ClientService{}

	//creating a new gRPC server
	grpc_server := grpc.NewServer()

	//Registering server and client services
	server.RegisterServerServiceServer(grpc_server, ServerService)
	client.RegisterClientServiceServer(grpc_server, ClientService)

	return grpc_server
}
