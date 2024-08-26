package main

import (
	"context"
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/server"
	"github.com/RouteHub-Link/routehub.client.hub/services"
)

func main() {
	ctx := context.Background()
	services.NewLoggerConfigurer(slog.LevelDebug)

	logger := services.GetLogger()
	cf := services.GetRedisConfig()

	logger.Log(ctx, slog.LevelDebug, "redis config", slog.String("host", cf.Host), slog.String("port", cf.Port), slog.String("password", cf.Password), slog.Int("db", cf.DB))
	_ = services.NewRedisClient(ctx, cf)

	server.Serve()
}
