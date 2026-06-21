package errors

import "errors"

var (
	ErrInvalidToken        = errors.New("invalid token")
	ErrUserWalletNotFound  = errors.New("user wallet not found")
	ErrDuplicateReference  = errors.New("transaction reference already exists")
	ErrInsufficientBalance = errors.New("your balance is insufficient for this transaction")
	ErrInternalServerError = errors.New("an unexpected errors occurred. Please try again later")
)
