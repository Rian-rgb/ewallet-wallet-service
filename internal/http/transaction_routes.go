package http

import (
	"ewallet-wallet/infra"
	"github.com/Rian-rgb/ewallet-common-lib/middleware"
	"github.com/gin-gonic/gin"
)

func registerTransactionRoutes(
	api *gin.RouterGroup,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	walletTransaction := api.Group("/wallet-transaction")
	walletTransaction.Use(
		middleware.AuthMiddleware(
			appDeps.JWTManager.ValidateToken,
			*appDeps.RedisRepo,
		))

	walletTransaction.PUT(
		"/credit",
		dependency.TransactionAPI.CreditBalance,
	)

	walletTransaction.PUT(
		"/debit",
		dependency.TransactionAPI.DebitBalance,
	)

	walletTransaction.GET(
		"/history",
		dependency.TransactionAPI.GetPaginition,
	)
}
