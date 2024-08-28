package server

import (
	"context"
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/packages"
	server_context "github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/router"
	"github.com/RouteHub-Link/routehub.client.hub/services"
	"github.com/labstack/echo/v4"
)

func NewRestServer() {
	ctx := context.Background()

	services.NewLoggerConfigurer(slog.LevelDebug)
	logger := services.GetLogger()

	packages.NewClientContainer(services.GetRedisClient(), logger, services.GetDetailsConfig())
	_ = services.NewTimeScaleClient(ctx, services.GetDetailsConfig().TimeScaleDB)

	e := echo.New()
	server_context.ApplyMiddleware(e)

	router.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
