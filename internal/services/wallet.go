package services

import (
	"context"
	"ewallet-wallet/internal/domain/wallet"
	internalErrors "ewallet-wallet/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WalletService struct {
	WalletRepo wallet.IRepository
}

func (s *WalletService) CreateWallet(ctx context.Context, walletEntity *wallet.Entity) (*wallet.Entity, error) {
	err := s.WalletRepo.Save(walletEntity)
	if err != nil {
		logger.WithContext(ctx).Error("failed to save walletEntity: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	return walletEntity, nil
}

func (s *WalletService) GetBalance(ctx context.Context, userID int) (balance float64, err error) {
	walletEntity, err := s.WalletRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, internalErrors.ErrUserWalletNotFound
		}

		logger.WithContext(ctx).Error("failed to find by user id: ", err)
		return 0, internalErrors.ErrInternalServerError
	}

	return walletEntity.Balance, nil
}
