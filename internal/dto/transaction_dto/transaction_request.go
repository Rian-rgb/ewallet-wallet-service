package transaction_dto

import (
	"ewallet-wallet/internal/domain/transaction"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/go-playground/validator/v10"
)

type TransactionRequest struct {
	Amount float64 `json:"amount" validate:"required" example:"10000"`
}

func (req TransactionRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *TransactionRequest) ToEntity() *transaction.Entity {
	return &transaction.Entity{
		Amount: req.Amount,
	}
}
