package http

import (
	"ewallet-wallet/infra"
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/gin-gonic/gin"
)

func NewServer(
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) *gin.Engine {
	router := gin.Default()

	RegisterRoutes(
		router,
		dependency,
		appDeps,
	)

	return router
}

func Serve(
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	router := NewServer(
		dependency,
		appDeps,
	)

	port := config.GetEnv("PORT", "8081")

	logger.Info("HTTP server listening on port: ", port)

	if err := router.Run(":" + port); err != nil {
		logger.Error("failed to start HTTP server: ", err)
	}
}
