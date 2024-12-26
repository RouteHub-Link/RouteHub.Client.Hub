package server

import (
	"log/slog"
	"os"
	"strings"

	"github.com/RouteHub-Link/routehub.client.hub/packages"
	server_context "github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/middlewares"
	"github.com/RouteHub-Link/routehub.client.hub/server/router"
	"github.com/RouteHub-Link/routehub.client.hub/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ory/graceful"
)

func NewRestServer() {

	logger := services.GetLogger()

	packages.NewMQTTClientContainer(services.GetRedisClient(), logger, services.GetDetailsConfig())

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
	}))

	server_context.ApplyMiddleware(e)

	router.ConfigureRoutes(e)

	e.Use(middleware.RateLimiterWithConfig(middlewares.RateConfig))

	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = "8080"
	}

	e.Server.Addr = strings.Join([]string{":", envPort}, "")

	server := graceful.WithDefaults(e.Server)

	logger.Info("main: Starting server")
	logger.Info("main: Listening on", slog.String("port", envPort))
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		logger.Error("main: Failed to gracefully shutdown", slog.Any("error", err))
	}
	logger.Info("main: Server stopped")
}
