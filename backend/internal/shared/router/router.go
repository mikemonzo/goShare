package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mikemonzo/goshare/internal/tenant/infrastructure/adapter/inbound"
)

func SetRouter(tenantHandler *inbound.TenantHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		tenants := api.Group("/tenants")
		{
			tenants.POST("/", tenantHandler.CreateTenant)
		}
	}

	return router
}
