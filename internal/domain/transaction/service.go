package transaction

import "context"

type IService interface {
	CreditBalance(
		ctx context.Context,
		userID int,
		transaction *Entity,
	) (newBalance float64, err error)

	DebitBalance(
		ctx context.Context,
		userID int,
		walletTransaction *Entity,
	) (newBalance float64, err error)

	GetPagination(
		ctx context.Context,
		userID int,
		page int,
		limit int,
		transactionType string,
	) (
		transactionsEntity []*Entity,
		totalPage int,
		totalData int,
		err error,
	)
}
