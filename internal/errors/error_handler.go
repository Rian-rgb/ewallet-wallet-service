package errors

import (
	"errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
)

func HandleServiceError(ctx *gin.Context, err error) {
	var (
		errCodeUnauthorized        = appErrors.ErrCodeUnauthorized
		errCodeWalletNotFound      = appErrors.ErrCodeWalletNotFound
		errCodeDuplicateReference  = appErrors.ErrCodeDuplicateReference
		errCodeInternalServerError = appErrors.ErrCodeInternalServerError
		errCodeInsufficientBalance = appErrors.ErrCodeInsufficientBalance
		errCodeUnknownError        = appErrors.ErrCodeUnknownError
	)
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrInvalidToken):

		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, err.Error())

	case errors.Is(err, ErrUserWalletNotFound):

		response.SendError(ctx, errCodeWalletNotFound.ToHTTPStatus(), errCodeWalletNotFound, err.Error())

	case errors.Is(err, ErrDuplicateReference):

		response.SendError(ctx, errCodeDuplicateReference.ToHTTPStatus(), errCodeDuplicateReference, err.Error())

	case errors.Is(err, ErrInsufficientBalance):
		response.SendError(ctx, errCodeInsufficientBalance.ToHTTPStatus(), errCodeInsufficientBalance, err.Error())

	case errors.Is(err, ErrInternalServerError):
		response.SendError(ctx, errCodeInternalServerError.ToHTTPStatus(), errCodeInternalServerError, err.Error())

	default:
		response.SendError(ctx, errCodeUnknownError.ToHTTPStatus(), errCodeUnknownError, response.InternalServerErrorMessage)
	}
}
