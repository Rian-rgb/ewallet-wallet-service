package handler

import (
	"ewallet-wallet/internal/domain/wallet"
	"ewallet-wallet/internal/dto/wallet_dto"
	"ewallet-wallet/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletHandler struct {
	WalletSvc wallet.IService
}

// @Summary		Create User's Wallet
// @Description	Registers a new wallet for the user.
// @Tags		Wallet
// @Accept		json
// @Produce		json
//
// @Param		Authorization	header		string								true	"Bearer <token>"
// @Param		request			body		wallet_dto.CreateWalletRequest		true	"Request Body"
//
// @Success		201	{object}	response.SuccessResponse{data=wallet_dto.CreateWalletResponse}	"Created"
// @Failure		400	{object}	response.BadRequestResponse										"Bad Request"
// @Failure		401	{object}	response.ErrorResponse											"Unauthorized"
// @Failure		500	{object}	response.ErrorResponse											"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/wallet/ [post]
func (hdl *WalletHandler) Create(ctx *gin.Context) {
	var (
		req            wallet_dto.CreateWalletRequest
		codeBadRequest = appErrors.ErrCodeBadRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	walletEntity, err := hdl.WalletSvc.CreateWallet(ctx, req.ToEntity())
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := &wallet_dto.CreateWalletResponse{UserID: walletEntity.UserID}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}

// @Summary		Get Balance of the User's Wallet
// @Description	Retrieves the current balance of the user's wallet.
// @Tags		Wallet
// @Accept		json
// @Produce		json
//
// @Param		Authorization	header		string		true	"Bearer <token>"
//
// @Success		200	{object}	response.SuccessResponse{data=wallet_dto.GetBalanceResponse}	"Success"
// @Failure		400	{object}	response.BadRequestResponse										"Bad Request"
// @Failure		401	{object}	response.ErrorResponse											"Unauthorized"
// @Failure		500	{object}	response.ErrorResponse											"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/wallet/balance [get]
func (hdl *WalletHandler) GetBalance(ctx *gin.Context) {

	errCodeUnauthorized := appErrors.ErrCodeUnauthorized

	userData, exists := security.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	balance, err := hdl.WalletSvc.GetBalance(ctx, userData.UserID)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := wallet_dto.GetBalanceResponse{Balance: balance}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
