package infra

import (
	"ewallet-wallet/infra/grpc"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"gorm.io/gorm"
)

type AppDependencies struct {
	PostgresDB   *gorm.DB
	RedisRepo    *redis.RedisRepository
	JWTManager   *security.JWTManager
	GrpcRegistry *grpc.ConnRegistry
}
