package infra

import (
	"ewallet-wallet/infra/grpc"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"gorm.io/gorm"
)

type AppDependencies struct {
	PostgresDB   *gorm.DB
	RedisRepo    *redis.RedisRepository
	GrpcRegistry *grpc.ConnRegistry
}
