package transaction

import "github.com/gin-gonic/gin"

type IHandler interface {
	CreditBalance(ctx *gin.Context)
	DebitBalance(ctx *gin.Context)
	GetPaginition(ctx *gin.Context)
}
