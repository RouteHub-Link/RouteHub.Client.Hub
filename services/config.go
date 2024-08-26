package services

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
	Redis   *RedisConfig
	Details *DetailsConfig
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

type DetailsConfig struct {
	OrganizationId string      `env:"ORGANIZATION_ID"`
	OwnerId        string      `env:"OWNER_ID"`
	PlatformId     string      `env:"PLATFORM_ID"`
	PlatformSecret string      `env:"PLATFORM_SECRET"`
	SEED           bool        `env:"SEED"`
	HostingMode    HostingMode `env:"HOSTING_MODE"`
}

type HostingMode string

const (
	HostingModeMQQT HostingMode = "MQQT"
	HostingModeRest HostingMode = "REST"
)

func getApplicationConfig() *ApplicationConfig {
	onceConfigure.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			logger.Log(context.Background(), slog.LevelError, "Error loading .env file", slog.String("error", err.Error()))
		}

		_appConfig = &ApplicationConfig{}
		_redisConfig := &RedisConfig{}
		_detailsConfig := &DetailsConfig{}

		env.Parse(_redisConfig)
		_appConfig.Redis = _redisConfig

		env.Parse(_detailsConfig)
		_appConfig.Details = _detailsConfig
	})

	return _appConfig
}

func GetHostingMode() HostingMode {
	appConfig := getApplicationConfig()
	if appConfig == nil {
		logger.Log(context.Background(), slog.LevelError, "Application config is nil")
		return HostingModeRest
	}

	return appConfig.Details.HostingMode
}

func GetRedisConfig() RedisConfig {
	appConfig := getApplicationConfig()
	if appConfig == nil {
		logger.Log(context.Background(), slog.LevelError, "Application config is nil")
		return RedisConfig{} // Return a default RedisConfig if appConfig is nil
	}

	return *appConfig.Redis
}

func GetDetailsConfig() DetailsConfig {
	appConfig := getApplicationConfig()
	if appConfig == nil {
		logger.Log(context.Background(), slog.LevelError, "Application config is nil")
	}

	return *appConfig.Details
}
