package transaction_dto

import (
	"ewallet-wallet/internal/domain/transaction"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/go-playground/validator/v10"
)

type TransactionRequest struct {
	Reference string  `json:"reference" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`
}

func (req TransactionRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *TransactionRequest) ToEntity() *transaction.Entity {
	return &transaction.Entity{
		Reference: req.Reference,
		Amount:    req.Amount,
	}
}
