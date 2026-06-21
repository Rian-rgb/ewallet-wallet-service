package wallet

import "context"

type IService interface {
	CreateWallet(ctx context.Context, walletEntity *Entity) (*Entity, error)
	GetBalance(ctx context.Context, userID int) (balance float64, err error)
}
