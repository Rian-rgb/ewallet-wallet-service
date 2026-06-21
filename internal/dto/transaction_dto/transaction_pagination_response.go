package transaction_dto

import (
	"time"
)

type TransactionPaginationResponse struct {
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}
