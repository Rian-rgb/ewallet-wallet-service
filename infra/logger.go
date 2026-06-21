package infra

import (
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
)

func InitLogger() {
	env := config.GetEnv("ENV", "production")
	logger.SetupLogger(env)
}
