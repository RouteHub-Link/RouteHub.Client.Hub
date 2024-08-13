package context

import (
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/clients"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (cc *CustomContext) GetLogger() *slog.Logger {
	return cc.Get("logger").(*slog.Logger)
}

func ApplyMiddleware(e *echo.Echo) {
	lc := clients.NewLoggerConfigurer(slog.LevelInfo)
	implementEchoLogger(e, lc)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})
}
