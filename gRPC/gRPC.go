package grpc

import (
	v1 "github.com/TheLazarusNetwork/erebrus/gRPC/v1"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {
	grpc_server := v1.Initialize()
	return grpc_server
}
