package main

import (
	"context"
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/mq"
	"github.com/RouteHub-Link/routehub.client.hub/packages"
	"github.com/RouteHub-Link/routehub.client.hub/server"
	"github.com/RouteHub-Link/routehub.client.hub/services"
)

func main() {
	ctx := context.Background()
	services.NewLoggerConfigurer(slog.LevelDebug)

	logger := services.GetLogger()
	cf := services.GetRedisConfig()

	logger.Log(ctx, slog.LevelDebug, "redis config", slog.String("host", cf.Host), slog.String("port", cf.Port), slog.String("password", cf.Password), slog.Int("db", cf.DB))
	rc := services.NewRedisClient(ctx, cf)

	packages.NewClientContainer(rc, logger, services.GetDetailsConfig())
	hostingMode := services.GetHostingMode()
	switch hostingMode {
	case services.HostingModeMQQT:
		logger.Log(ctx, slog.LevelDebug, "MQQT Hosting Mode")
		mq.NewMQQTServer()
	case services.HostingModeRest:
		logger.Log(ctx, slog.LevelDebug, "Rest Hosting Mode")
		server.Serve()
	}

}
