package wallet

import "github.com/gin-gonic/gin"

type IHandler interface {
	Create(ctx *gin.Context)
	GetBalance(ctx *gin.Context)
}
