package wallet

import "gorm.io/gorm"

type IRepository interface {
	Save(wallet *Entity) error
	FindByUserID(userID int) (walletEntity *Entity, err error)
	UpdateBalanceTx(tx *gorm.DB, walletID int, newBalance float64) error
}
