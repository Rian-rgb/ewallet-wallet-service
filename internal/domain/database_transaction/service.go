package database_transaction

import "gorm.io/gorm"

type IService interface {
	Transaction(func(tx *gorm.DB) error) error
}
