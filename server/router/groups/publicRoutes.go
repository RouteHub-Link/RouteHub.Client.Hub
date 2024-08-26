package groups

import (
	"github.com/RouteHub-Link/routehub.client.hub/server/router/handlers"
	"github.com/labstack/echo/v4"
)

func MapPublicRoutes(e *echo.Echo) {
	echoHandlers := handlers.NewEchoHandlers()

	e.GET("/", echoHandlers.HomeHandler)
	e.GET("/pins", echoHandlers.PinHandler)
	e.GET("/:key", echoHandlers.HandleShortenURL)
}
