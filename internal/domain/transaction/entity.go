package transaction

import "time"

type Entity struct {
	ID              int     `gorm:"primaryKey"`
	WalletID        int     `gorm:"column:wallet_id"`
	Amount          float64 `gorm:"column:amount;type:decimal(15,2)"`
	TransactionType Type    `gorm:"column:transaction_type;type:wallet_transaction_type"`
	Reference       string  `gorm:"column:reference;type:varchar(100);unique"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (*Entity) TableName() string {
	return "wallet_transactions"
}
