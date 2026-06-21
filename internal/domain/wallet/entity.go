package wallet

import "time"

type Entity struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"column:user_id;unique"`
	Balance   float64 `gorm:"column:balance;type:decimal(15, 2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Entity) TableName() string {
	return "wallets"
}
