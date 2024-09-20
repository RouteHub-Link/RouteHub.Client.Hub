package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RouteHub-Link/routehub.client.hub/packages"
	"github.com/RouteHub-Link/routehub.client.hub/packages/analytics"
	server_context "github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/middlewares"
	"github.com/RouteHub-Link/routehub.client.hub/server/router"
	"github.com/RouteHub-Link/routehub.client.hub/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRestServer() {
	ctx := context.Background()

	logger := services.GetLogger()

	packages.NewMQTTClientContainer(services.GetRedisClient(), logger, services.GetDetailsConfig())
	dbpool := services.NewTimeScaleClient(ctx, services.GetDetailsConfig().TimeScaleDB)
	defer dbpool.Close()

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

	analyticsMiddleware, err := analytics.NewAnalyticsMiddleware(
		dbpool,
		analytics.WithLogger(logger),
		analytics.WithMaxHeaderValueLength(4096),
	)

	if err != nil {
		logger.Error("Failed to create analytics middleware", "error", err)
		os.Exit(1)
	}

	e.Use(analyticsMiddleware.Middleware())

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server gracefully", "error", err)
	}
}
