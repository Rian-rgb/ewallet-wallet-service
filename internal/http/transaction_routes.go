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
			dependency.UmsClient.ValidateToken,
			*appDeps.RedisRepo,
		))

	walletTransaction.PUT(
		"/credit",
		dependency.TransactionHdl.CreditBalance,
	)

	walletTransaction.PUT(
		"/debit",
		dependency.TransactionHdl.DebitBalance,
	)

	walletTransaction.GET(
		"/history",
		dependency.TransactionHdl.GetPaginition,
	)
}
