package infra

import (
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
)

func InitRedis() *redis.RedisRepository {
	redisHost := config.GetEnv("REDIS_HOST", "")
	redisPassword := config.GetEnv("REDIS_PASSWORD", "")

	rdb, err := redis.NewClient(redisHost, redisPassword, 0)
	if err != nil {
		logger.Error("failed connect to redis: %v", err)
	}

	redisRepo := &redis.RedisRepository{Core: rdb}

	return redisRepo
}
