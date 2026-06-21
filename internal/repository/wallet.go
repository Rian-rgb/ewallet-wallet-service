package repository

import (
	"ewallet-wallet/internal/domain/wallet"
	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func (repo *WalletRepository) Save(wallet *wallet.Entity) error {
	return repo.DB.Create(wallet).Error
}

func (repo *WalletRepository) FindByUserID(userID int) (walletEntity *wallet.Entity, err error) {
	err = repo.DB.Where("user_id = ?", userID).First(&walletEntity).Error

	return walletEntity, err
}

func (r *WalletRepository) UpdateBalanceTx(tx *gorm.DB, walletID int, newBalance float64) error {
	return tx.
		Model(&wallet.Entity{}).
		Where("id = ?", walletID).
		Update("balance", newBalance).
		Error
}
