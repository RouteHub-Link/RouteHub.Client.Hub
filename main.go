package main

import (
	"context"
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/mq"
	"github.com/RouteHub-Link/routehub.client.hub/server"
	"github.com/RouteHub-Link/routehub.client.hub/services"
)

func main() {
	ctx := context.Background()
	services.NewLoggerConfigurer(slog.LevelDebug)
	logger := services.GetLogger()

	hostingMode := services.GetHostingMode()

	cf := services.GetRedisConfig()
	logger.Log(ctx, slog.LevelDebug, "redis config", slog.String("host", cf.Host), slog.String("port", cf.Port), slog.String("password", cf.Password), slog.Int("db", cf.DB))

	rc := services.NewRedisClient(ctx, cf)

	switch hostingMode {
	case services.HostingModeRest:
		logger.Log(ctx, slog.LevelDebug, "Rest Hosting Mode")
		server.NewRestServer()
	case services.HostingModeMQTT:
		logger.Log(ctx, slog.LevelDebug, "MQTT Hosting Mode")
		mq.NewMQTTServer(rc)
	default:
		logger.Log(ctx, slog.LevelError, "Invalid hosting mode")
	}
}
