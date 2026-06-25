package infra

import (
	"ewallet-wallet/external/ums"
	"ewallet-wallet/internal/domain/transaction"
	"ewallet-wallet/internal/domain/wallet"
	"ewallet-wallet/internal/handler"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/service"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type Dependency struct {
	WalletRepo     wallet.IRepository
	WalletHdl      wallet.IHandler
	TransactionHdl transaction.IHandler
	UmsClient      *ums.Client
}

func DependencyInject(appDeps *AppDependencies) *Dependency {

	walletRepo := &repository.WalletRepository{
		DB: appDeps.PostgresDB,
	}

	walletTransactionRepo := &repository.TransactionRepository{
		DB: appDeps.PostgresDB,
	}

	uow := &service.UnitOfWork{
		DB: appDeps.PostgresDB,
	}
	walletSvc := &service.WalletService{
		WalletRepo: walletRepo,
	}
	walletTransactionSvc := &service.TransactionService{
		TransactionRepo: walletTransactionRepo,
		WalletRepo:      walletRepo,
		Uow:             uow,
	}

	walletHdl := &handler.WalletHandler{
		WalletSvc: walletSvc,
	}
	walletTransactionHdl := &handler.TransactionHandler{
		TransactionSvc: walletTransactionSvc,
	}

	pbUmsClient := pb.NewTokenValidationServiceClient(appDeps.GrpcRegistry.UmsConn.Conn)
	umsGrpcClient := ums.NewClient(pbUmsClient)

	return &Dependency{
		WalletRepo:     walletRepo,
		WalletHdl:      walletHdl,
		TransactionHdl: walletTransactionHdl,
		UmsClient:      umsGrpcClient,
	}
}
