package router

import (
	"github.com/RouteHub-Link/routehub.client.hub/server/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(e *echo.Echo) {
	e.Static("/static", "./public")
	e.Pre(middleware.RemoveTrailingSlash())

	groups.MapPublicRoutes(e)
	groups.MapMiscRoutes(e)
}
