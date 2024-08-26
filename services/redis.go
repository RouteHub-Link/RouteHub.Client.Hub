package services

import (
	"context"
	"log/slog"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	onceRedisClient sync.Once
	redisClient     *redis.Client
)

func NewRedisClient(ctx context.Context, cr RedisConfig) *redis.Client {
	onceRedisClient.Do(func() {
		addr := strings.Join([]string{cr.Host, cr.Port}, ":")

		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: cr.Password,
			DB:       cr.DB,
		})

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			logger.Log(ctx, slog.LevelError, "Error pinging redis", slog.String("error", err.Error()))
		} else {
			logger.Log(ctx, slog.LevelInfo, "Redis connected", slog.String("addr", addr))
		}

	})

	return redisClient
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		logger.Log(context.Background(), slog.LevelError, "Redis client is nil")
	}

	return redisClient
}
