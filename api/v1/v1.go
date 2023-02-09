package v1

import (
	"github.com/TheLazarusNetwork/erebrus/api/v1/authenticate"
	"github.com/TheLazarusNetwork/erebrus/api/v1/client"
	"github.com/TheLazarusNetwork/erebrus/api/v1/server"
	"github.com/TheLazarusNetwork/erebrus/api/v1/status"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Setup API EndPoints
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		client.ApplyRoutes(v1)
		server.ApplyRoutes(v1)
		status.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)

	}
}
