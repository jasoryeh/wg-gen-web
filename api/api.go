package api

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/jasoryeh/wg-gen-web/api/v1"
)

// ApplyRoutes apply routes to gin engine
func ApplyRoutes(r *gin.Engine, private bool) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api, private)
	}
}
