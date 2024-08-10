package groups

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
)

var (
	mockMeta = components.MetaDescription{Title: "RouteHub", Description: "RouteHub is a platform that allows you to create, share, and discover routes for your favorite activities."}
)

func MapPublicRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
}

func HomeHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, pages.Home(mockMeta))
}
