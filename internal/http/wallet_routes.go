package http

import (
	"ewallet-wallet/infra"
	"github.com/Rian-rgb/ewallet-common-lib/middleware"
	"github.com/gin-gonic/gin"
)

func registerWalletRoutes(
	api *gin.RouterGroup,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	wallet := api.Group("/wallet")

	wallet.POST(
		"",
		dependency.WalletHdl.Create,
	)

	walletAuth := wallet
	walletAuth.Use(
		middleware.AuthMiddleware(
			dependency.UmsClient.ValidateToken,
			*appDeps.RedisRepo,
		))

	walletAuth.GET(
		"/balance",
		dependency.WalletHdl.GetBalance,
	)
}
