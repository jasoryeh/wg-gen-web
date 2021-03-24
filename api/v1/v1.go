package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasoryeh/wg-gen-web/api/v1/auth"
	"github.com/jasoryeh/wg-gen-web/api/v1/client"
	"github.com/jasoryeh/wg-gen-web/api/v1/server"
	"github.com/jasoryeh/wg-gen-web/api/v1/status"
)

// ApplyRoutes apply routes to gin router
func ApplyRoutes(r *gin.RouterGroup, private bool) {
	v1 := r.Group("/v1.0")
	{
		if private {
			client.ApplyRoutes(v1)
			server.ApplyRoutes(v1)
			status.ApplyRoutes(v1)
		} else {
			auth.ApplyRoutes(v1)
		}
	}
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello",
	})
}
