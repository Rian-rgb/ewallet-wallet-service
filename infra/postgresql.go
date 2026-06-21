package infra

import (
	"ewallet-wallet/internal/domain/transaction"
	"ewallet-wallet/internal/domain/wallet"
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/database"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"gorm.io/gorm"
)

func InitPostgresql() *gorm.DB {
	dbCfg := database.PostgresConfig{
		Host:         config.GetEnv("DB_HOST", ""),
		User:         config.GetEnv("DB_USER", ""),
		Password:     config.GetEnv("DB_PASSWORD", ""),
		DBName:       config.GetEnv("DB_NAME", "ewallet_ums"),
		Port:         config.GetEnv("DB_PORT", "8080"),
		MaxIdleConns: 5,
		MaxOpenConns: 20,
	}

	dbClient, err := database.NewPostgresClient(dbCfg)
	if err != nil {
		logger.Error("the App shutdown because failed connect DB: ", err)
	}

	err = dbClient.AutoMigrate(&wallet.Entity{}, &transaction.Entity{})
	if err != nil {
		logger.Error("failed doing auto migrate to internal tabel: ", err)
	}

	return dbClient
}
