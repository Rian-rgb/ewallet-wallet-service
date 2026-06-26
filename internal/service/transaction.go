package service

import (
	"context"
	"ewallet-wallet/internal/domain/database_transaction"
	"ewallet-wallet/internal/domain/transaction"
	"ewallet-wallet/internal/domain/wallet"
	internalErrors "ewallet-wallet/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math"
)

type TransactionService struct {
	Uow             database_transaction.IService
	TransactionRepo transaction.IRepository
	WalletRepo      wallet.IRepository
}

func (svc *TransactionService) CreditBalance(
	ctx context.Context,
	userID int,
	walletTransaction *transaction.Entity,
) (newBalance float64, err error) {

	exists, err := svc.TransactionRepo.IsReferenceExists(walletTransaction.Reference)
	if err != nil {

		logger.WithContext(ctx).Error("failed to check walletTransaction reference existence: ", err)
		return 0, internalErrors.ErrInternalServerError
	}

	if exists {
		logger.WithContext(ctx).Error("duplicate walletTransaction reference")
		return 0, internalErrors.ErrDuplicateReference
	}

	walletEntity, err := svc.WalletRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, internalErrors.ErrUserWalletNotFound
		}

		logger.WithContext(ctx).Error("failed to find by user id: ", err)
		return 0, internalErrors.ErrInternalServerError
	}

	err = svc.Uow.Transaction(func(tx *gorm.DB) error {
		newBalance = walletEntity.Balance + walletTransaction.Amount

		if err := svc.WalletRepo.UpdateBalanceTx(
			tx,
			walletEntity.ID,
			newBalance,
		); err != nil {

			logger.WithContext(ctx).Error("failed to update balance: ", err)
			return internalErrors.ErrInternalServerError
		}

		walletTransaction.WalletID = walletEntity.ID
		walletTransaction.Reference = utils.GenerateUUID()
		walletTransaction.TransactionType = transaction.Credit

		if err := svc.TransactionRepo.SaveTx(
			tx,
			walletTransaction,
		); err != nil {

			logger.WithContext(ctx).Error("failed to save transaction: ", err)
			return internalErrors.ErrInternalServerError
		}

		return nil
	})

	if err != nil {
		return 0, internalErrors.ErrInternalServerError
	}

	return newBalance, nil
}

func (svc *TransactionService) DebitBalance(
	ctx context.Context,
	userID int,
	walletTransaction *transaction.Entity,
) (newBalance float64, err error) {

	exists, err := svc.TransactionRepo.IsReferenceExists(walletTransaction.Reference)
	if err != nil {

		logger.WithContext(ctx).Error("failed to check walletTransaction reference existence: ", err)
		return 0, internalErrors.ErrInternalServerError
	}

	if exists {
		logger.WithContext(ctx).Error("duplicate walletTransaction reference")
		return 0, internalErrors.ErrDuplicateReference
	}

	walletEntity, err := svc.WalletRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, internalErrors.ErrUserWalletNotFound
		}

		logger.WithContext(ctx).Error("failed to find by user id: ", err)
		return 0, internalErrors.ErrInternalServerError
	}

	newBalance = walletEntity.Balance - walletTransaction.Amount
	if newBalance < 0 {
		return 0, internalErrors.ErrInsufficientBalance
	}

	err = svc.Uow.Transaction(func(tx *gorm.DB) error {

		if err := svc.WalletRepo.UpdateBalanceTx(
			tx,
			walletEntity.ID,
			newBalance,
		); err != nil {

			logger.WithContext(ctx).Error("failed to update balance: ", err)
			return internalErrors.ErrInternalServerError
		}

		walletTransaction.WalletID = walletEntity.ID
		walletTransaction.Reference = utils.GenerateUUID()
		walletTransaction.TransactionType = transaction.Debit

		if err := svc.TransactionRepo.SaveTx(
			tx,
			walletTransaction,
		); err != nil {

			logger.WithContext(ctx).Error("failed to save transaction: ", err)
			return internalErrors.ErrInternalServerError
		}

		return nil
	})

	if err != nil {
		return 0, internalErrors.ErrInternalServerError
	}

	return newBalance, nil
}

func (svc *TransactionService) GetPagination(
	ctx context.Context,
	userID int,
	page int,
	limit int,
	transactionType string,
) (
	transactionsEntity []*transaction.Entity,
	totalPage int,
	totalData int,
	err error,
) {

	walletEntity, err := svc.WalletRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, 0, internalErrors.ErrUserWalletNotFound
		}

		logger.WithContext(ctx).Error("failed to find by user id: ", err)
		return nil, 0, 0, internalErrors.ErrInternalServerError
	}

	offset := (page - 1) * limit
	transactionsEntity, err = svc.TransactionRepo.FindPagination(walletEntity.ID, offset, limit, transactionType)
	if err != nil {
		logger.WithContext(ctx).Error("failed to find transacton history: ", err)
		return nil, 0, 0, internalErrors.ErrInternalServerError
	}

	resultTotalData, err := svc.TransactionRepo.Count(walletEntity.ID, transactionType)
	if err != nil {
		logger.WithContext(ctx).Error("failed to count transaction history: ", err)
		return nil, 0, 0, internalErrors.ErrInternalServerError
	}

	totalData = int(resultTotalData)
	totalPage = int(
		math.Ceil(
			float64(totalData) / float64(limit),
		),
	)

	return transactionsEntity, totalPage, totalData, nil
}
