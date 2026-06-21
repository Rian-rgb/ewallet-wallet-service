package infra

import (
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

func InitJWT() *security.JWTManager {
	jwtSecret := config.GetEnv("APP_SECRET", "")
	appName := config.GetEnv("APP_NAME", "")

	jwtManager := security.NewJWTManager(jwtSecret, appName)

	return jwtManager
}
