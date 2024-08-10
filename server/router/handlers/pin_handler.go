package handlers

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
)

func (eh echoHandlers) PinHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, pages.Pins(eh.layoutDescription))
}
