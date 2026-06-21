package transaction_dto

import (
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/go-playground/validator/v10"
)

type TransactionPaginationRequest struct {
	Page            int    `form:"page" validate:"gte=1"`
	Size            int    `form:"size" validate:"gte=1,lte=100"`
	TransactionType string `form:"transactionType" validate:"omitempty,oneof=CREDIT DEBIT"`
}

func (req TransactionPaginationRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}
