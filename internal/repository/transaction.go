package repository

import (
	"ewallet-wallet/internal/domain/transaction"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func (r *TransactionRepository) SaveTx(tx *gorm.DB, entity *transaction.Entity) error {
	return tx.Create(entity).Error
}

func (repo *TransactionRepository) IsReferenceExists(reference string) (bool, error) {
	var count int64

	err := repo.DB.
		Model(&transaction.Entity{}).
		Where("reference = ?", reference).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *TransactionRepository) FindPagination(
	walletID int,
	offset int,
	limit int,
	transactionType string,
) (transactionsEntity []*transaction.Entity, err error) {

	query := repo.DB.
		Where("wallet_id = ?", walletID)

	if transactionType != "" {
		query = query.Where("transaction_type = ?", transactionType)
	}

	err = query.
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactionsEntity).
		Error

	return transactionsEntity, err
}

func (repo *TransactionRepository) Count(walletID int, transactionType string) (total int64, err error) {

	query := repo.DB.
		Model(&transaction.Entity{}).
		Where("wallet_id = ?", walletID)

	if transactionType != "" {
		query = query.Where(
			"transaction_type = ?",
			transactionType,
		)
	}

	err = query.Count(&total).Error

	return total, err
}
