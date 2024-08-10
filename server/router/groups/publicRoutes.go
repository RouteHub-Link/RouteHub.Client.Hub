package groups

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func MapPublicRoutes(e *echo.Echo) {
	http.Handle("/", templ.Handler(pages.Home()))

	e.GET("/", HomeHandler)
}

func HomeHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, pages.Home())
}
