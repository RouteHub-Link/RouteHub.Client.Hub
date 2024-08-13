package main

import (
	"context"
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/clients"
	"github.com/RouteHub-Link/routehub.client.hub/server"
)

func main() {
	ctx := context.Background()
	clients.NewLoggerConfigurer(slog.LevelDebug)

	logger := clients.GetLogger()
	cf := clients.GetRedisConfig()
	logger.Log(ctx, slog.LevelDebug, "redis config", slog.String("host", cf.Host), slog.String("port", cf.Port), slog.String("password", cf.Password), slog.Int("db", cf.DB))
	_ = clients.NewRedisClient(ctx, cf)

	server.Serve()
}
