package http

import (
	"ewallet-wallet/infra"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	api := router.Group("/api/v1")

	registerWalletRoutes(
		api,
		dependency,
		appDeps,
	)

	registerTransactionRoutes(
		api,
		dependency,
		appDeps,
	)

	registerSwaggerRoutes(router)
}
