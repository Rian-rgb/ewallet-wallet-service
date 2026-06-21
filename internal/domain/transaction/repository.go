package transaction

import "gorm.io/gorm"

type IRepository interface {
	SaveTx(tx *gorm.DB, entity *Entity) error
	IsReferenceExists(reference string) (bool, error)
	FindPagination(
		walletID int,
		offset int,
		limit int,
		transactionType string,
	) (transactionsEntity []*Entity, err error)
	Count(walletID int, transactionType string) (total int64, err error)
}
