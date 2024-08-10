package router

import (
	"github.com/RouteHub-Link/routehub.client.hub/server/router/groups"
	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	groups.MapPublicRoutes(e)
	groups.MapMiscRoutes(e)

	e.Static("/", "./public")
}
