package wallet_dto

import (
	"ewallet-wallet/internal/domain/wallet"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/go-playground/validator/v10"
)

type CreateWalletRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

func (req CreateWalletRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *CreateWalletRequest) ToEntity() *wallet.Entity {
	return &wallet.Entity{
		UserID: req.UserID,
	}
}
