package handler

import (
	"ewallet-wallet/internal/domain/wallet"
	"ewallet-wallet/internal/dto/wallet_dto"
	"ewallet-wallet/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	contextUtil "github.com/Rian-rgb/ewallet-common-lib/pkg/context"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletHandler struct {
	WalletSvc wallet.IService
}

// @Summary Create Wallet User
// @Description Registers a new wallet for the user. Note: Each user is limited to only one active wallet. This endpoint will return an error if a wallet already exists.
// @Accept       json
// @Produce      json
// @Param        request  body      wallet_dto.CreateWalletRequest  true  "Payload creates wallet user"
// @Success      200      {object}  response.SuccessResponse{data=wallet_dto.CreateWalletResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /wallet/ [post]
func (api *WalletHandler) Create(ctx *gin.Context) {
	var (
		req            wallet_dto.CreateWalletRequest
		codeBadRequest = appErrors.ErrCodeBadRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	walletEntity, err := api.WalletSvc.CreateWallet(ctx, req.ToEntity())
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := &wallet_dto.CreateWalletResponse{UserID: walletEntity.UserID}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}

// @Summary Get Wallet Balance User
// @Description Retrieves the current balance of the user's wallet. This endpoint returns the available funds with the authenticated user account. Note: This action does not modify any wallet data and provides the most up-to-date balance information.
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Success      200      {object}  response.SuccessResponse{data=wallet_dto.GetBalanceResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /wallet/balance [get]
func (api *WalletHandler) GetBalance(ctx *gin.Context) {

	errCodeUnauthorized := appErrors.ErrCodeUnauthorized

	userData, exists := contextUtil.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	balance, err := api.WalletSvc.GetBalance(ctx, userData.UserID)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := wallet_dto.GetBalanceResponse{Balance: balance}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
