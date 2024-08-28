package handlers

import (
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type MQHandler interface {
	Set(payload []byte) error
	Delete(payload []byte) error
	Get(payload []byte) ([]byte, error)
	Fetch(payload []byte) ([]byte, error)
}

type LinkHandlers struct {
	redisClient *redis.Client
	logger      *slog.Logger
}

func NewLinkHandlers(redisClient *redis.Client, logger *slog.Logger) MQHandler {
	return MQHandler(&LinkHandlers{
		redisClient: redisClient,
		logger:      logger,
	})
}

type PlatformHandlers struct {
	redisClient *redis.Client
	logger      *slog.Logger
}

func NewPlatformHandlers(redisClient *redis.Client, logger *slog.Logger) MQHandler {
	return MQHandler(&PlatformHandlers{
		redisClient: redisClient,
		logger:      logger,
	})
}
