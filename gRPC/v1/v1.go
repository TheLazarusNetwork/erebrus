package v1

import (
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/paseto"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/authenticate/selector"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/client"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/server"
	"github.com/TheLazarusNetwork/erebrus/gRPC/v1/status"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	selector_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {

	//get the instance of server and client services
	ServerService := &server.ServerService{}
	ClientService := &client.ClientService{}
	StatusService := &status.StatusService{}

	//creating a new gRPC server
	grpc_server := grpc.NewServer(
		grpc.ChainStreamInterceptor(selector_middleware.StreamServerInterceptor(
			auth.StreamServerInterceptor(paseto.PASETO), selector_middleware.MatchFunc(selector.LoginSkip))),
		grpc.ChainUnaryInterceptor(selector_middleware.UnaryServerInterceptor(
			auth.UnaryServerInterceptor(paseto.PASETO), selector_middleware.MatchFunc(selector.LoginSkip))),
	)
	server.RegisterServerServiceServer(grpc_server, ServerService)
	client.RegisterClientServiceServer(grpc_server, ClientService)
	status.RegisterStatusServiceServer(grpc_server, StatusService)

	return grpc_server
}
