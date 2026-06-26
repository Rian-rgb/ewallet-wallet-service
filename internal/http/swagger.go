package http

import (
	_ "ewallet-wallet/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           E Wallet API (Wallet Service)
// @version         0.0
// @description     API Service for wallet user.
// @description     <br/><b>Developer:</b> Muhammad Brilian Satria Utama
// @description     <b>Environment:</b> Development
// @host            localhost:8081
// @BasePath        /api/v1
func registerSwaggerRoutes(router *gin.Engine) {
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}
