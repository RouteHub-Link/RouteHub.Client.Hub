package handlers

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
)

func (eh echoHandlers) HomeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	sec := c.(*context.ServerEchoContext)
	platformClient := sec.GetPlatformClientService()
	platform, err := platformClient.GetPlatform(ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return extensions.Render(c, http.StatusOK, pages.Home(*platform.LayoutDescription))
}
