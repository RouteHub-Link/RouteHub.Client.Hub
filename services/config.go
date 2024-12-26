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
	OrganizationId string `env:"ORGANIZATION_ID"`
	OwnerId        string `env:"OWNER_ID"`
	PlatformId     string `env:"PLATFORM_ID"`
	PlatformSecret string `env:"PLATFORM_SECRET"`

	Name    string `env:"NAME"`
	Version string `env:"VERSION"`

	SEED        bool        `env:"SEED"`
	HostingMode HostingMode `env:"HOSTING_MODE"`
}

type HostingMode string

const (
	HostingModeMQTT HostingMode = "MQTT"
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
		logger.Log(context.Background(), slog.LevelDebug, "Redis config", slog.Any("config", _redisConfig))

		env.Parse(_detailsConfig)
		_appConfig.Details = _detailsConfig
		logger.Log(context.Background(), slog.LevelDebug, "Details config", slog.Any("config", _detailsConfig))
	})

	return _appConfig
}

func GetHostingMode() HostingMode {
	appConfig := getApplicationConfig()
	if appConfig == nil {
		logger.Log(context.Background(), slog.LevelError, "Application config is nil")
		return HostingModeRest
	}

	if appConfig.Details.HostingMode == "" {
		logger.Log(context.Background(), slog.LevelWarn, "Hosting mode is empty overrided", slog.Any("mode", HostingModeRest))
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
