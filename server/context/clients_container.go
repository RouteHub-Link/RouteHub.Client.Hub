package context

import (
	"github.com/RouteHub-Link/routehub.client.hub/packages"
	"github.com/labstack/echo/v4"
)

func implementClientsContainer(e *echo.Echo, cc *packages.ClientContainer, contextKey string) {
	// add loger to echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(contextKey, cc)
			return next(c)
		}
	})
}
