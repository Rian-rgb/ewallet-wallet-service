package services

import "gorm.io/gorm"

type UnitOfWork struct {
	DB *gorm.DB
}

func (u *UnitOfWork) Transaction(
	fn func(tx *gorm.DB) error,
) error {
	return u.DB.Transaction(fn)
}
