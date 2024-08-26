package clients

import (
	"context"
	"log/slog"
	"sync"

	"github.com/caarlos0/env"

	"github.com/joho/godotenv"
)

var (
	_appConfig    *ApplicationConfig
	onceConfigure sync.Once
)

type ApplicationConfig struct {
	Redis *RedisConfig
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

func getApplicationConfig() *ApplicationConfig {
	onceConfigure.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			logger.Log(context.Background(), slog.LevelError, "Error loading .env file", slog.String("error", err.Error()))
		}

		_appConfig = &ApplicationConfig{}
		_redisConfig := &RedisConfig{}

		env.Parse(_redisConfig)
		_appConfig.Redis = _redisConfig
	})

	return _appConfig
}

func GetRedisConfig() RedisConfig {
	appConfig := getApplicationConfig()
	if appConfig == nil {
		logger.Log(context.Background(), slog.LevelError, "Application config is nil")
	}

	return *appConfig.Redis
}
