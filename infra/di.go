package infra

import (
	"ewallet-wallet/external/ums"
	"ewallet-wallet/internal/domain/transaction"
	"ewallet-wallet/internal/domain/wallet"
	"ewallet-wallet/internal/handler"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type Dependency struct {
	WalletRepo     wallet.IRepository
	WalletAPI      wallet.IHandler
	TransactionAPI transaction.IHandler
	UmsClient      *ums.Client
}

func DependencyInject(appDeps *AppDependencies) *Dependency {

	walletRepo := &repository.WalletRepository{
		DB: appDeps.PostgresDB,
	}

	walletTransactionRepo := &repository.TransactionRepository{
		DB: appDeps.PostgresDB,
	}

	uow := &services.UnitOfWork{
		DB: appDeps.PostgresDB,
	}
	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletTransactionSvc := &services.TransactionService{
		TransactionRepo: walletTransactionRepo,
		WalletRepo:      walletRepo,
		Uow:             uow,
	}

	walletAPI := &handler.WalletHandler{
		WalletSvc: walletSvc,
	}
	walletTransactionAPI := &handler.TransactionHandler{
		TransactionSvc: walletTransactionSvc,
	}

	pbUmsClient := pb.NewTokenValidationServiceClient(appDeps.GrpcRegistry.UmsConn.Conn)
	umsGrpcClient := ums.NewClient(pbUmsClient)

	return &Dependency{
		WalletRepo:     walletRepo,
		WalletAPI:      walletAPI,
		TransactionAPI: walletTransactionAPI,
		UmsClient:      umsGrpcClient,
	}
}
