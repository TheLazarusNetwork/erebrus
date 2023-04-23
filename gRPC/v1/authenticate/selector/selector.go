package selector

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
)

func LoginSkip(_ context.Context, c interceptors.CallMeta) bool {
	methods := []string{"server.ServerService", "client.ClientService"}
	for _, s := range methods {
		if c.Service == s {
			return true
		}
	}
	return false
}
