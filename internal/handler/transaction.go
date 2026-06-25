package handler

import (
	"ewallet-wallet/internal/domain/transaction"
	"ewallet-wallet/internal/dto/transaction_dto"
	"ewallet-wallet/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	TransactionSvc transaction.IService
}

// @Summary Credit Wallet
// @Description Adds funds to the user's wallet.
// @Description Note: This endpoint updates the current wallet balance by applying the specified credit amount.
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Param        request  body      transaction_dto.TransactionRequest  true  "Payload credit wallet"
// @Success      200      {object}  response.SuccessResponse{data=transaction_dto.TransactionResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /wallet-transaction/credit [put]
func (hdl *TransactionHandler) CreditBalance(ctx *gin.Context) {
	var (
		req                 transaction_dto.TransactionRequest
		errCodeBadRequest   = appErrors.ErrCodeBadRequest
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.WithContext(ctx).Warn("request body validation failed")
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	userData, exists := security.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	transactionEntity := req.ToEntity()
	newBalance, err := hdl.TransactionSvc.CreditBalance(ctx, userData.UserID, transactionEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := transaction_dto.TransactionResponse{Balance: newBalance}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}

// @Summary Debit Wallet
// @Description Deducts funds from the user's wallet.
// @Description Note: This endpoint updates the wallet balance by applying the specified debit amount.
// @Description Note: This action will return an error if the current balance is insufficient.
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Param        request  body      transaction_dto.TransactionRequest  true  "Payload debit wallet"
// @Success      200      {object}  response.SuccessResponse{data=transaction_dto.TransactionResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /wallet-transaction/debit [put]
func (hdl *TransactionHandler) DebitBalance(ctx *gin.Context) {
	var (
		req                 transaction_dto.TransactionRequest
		errCodeBadRequest   = appErrors.ErrCodeBadRequest
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.WithContext(ctx).Warn("request body validation failed")
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	userData, exists := security.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	transactionEntity := req.ToEntity()
	newBalance, err := hdl.TransactionSvc.DebitBalance(ctx, userData.UserID, transactionEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := transaction_dto.TransactionResponse{Balance: newBalance}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}

// @Summary Get Pagination Transaction User
// @Description Retrieves the transaction history for the user's wallet.
// @Description Note: This endpoint returns a paginated list of credit and debit transactions.
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Param        page           query     int     false  "Page number"
// @Param        size          query     int     false  "Number of items per page"
// @Param        transactionType query string false "Filter by type (CREDIT/DEBIT)"
// @Success      200      {object}  response.SuccessResponse{data=[]transaction_dto.TransactionPaginationResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /wallet-transaction/history [get]
func (hdl *TransactionHandler) GetPaginition(ctx *gin.Context) {
	var (
		req                 transaction_dto.TransactionPaginationRequest
		errCodeBadRequest   = appErrors.ErrCodeBadRequest
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
	)

	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.WithContext(ctx).Error("invalid request parameter: ", err)
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidRequestMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.WithContext(ctx).Warn("request parameter validation failed")
		response.SendBadRequest(ctx, errCodeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	userData, exists := security.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	transactionsEntity, totalPage, totalData, err := hdl.TransactionSvc.GetPagination(
		ctx,
		userData.UserID,
		req.Page,
		req.Size,
		req.TransactionType,
	)

	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	respDatas := make(
		[]transaction_dto.TransactionPaginationResponse,
		0,
		len(transactionsEntity),
	)

	for _, transactionEntity := range transactionsEntity {
		respDatas = append(
			respDatas,
			transaction_dto.TransactionPaginationResponse{
				Amount:          transactionEntity.Amount,
				TransactionType: string(transactionEntity.TransactionType),
				CreatedAt:       transactionEntity.CreatedAt,
			},
		)
	}

	respMeta := response.PaginationMeta{
		CurrentPage: req.Page,
		PageSize:    req.Size,
		TotalPage:   totalPage,
		TotalData:   totalData,
	}

	response.SendSuccessWithMeta(ctx, http.StatusOK, response.SuccessMessage, respDatas, respMeta)
}
