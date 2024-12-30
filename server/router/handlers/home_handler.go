package handlers

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
	"github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	redirection_pages "github.com/RouteHub-Link/routehub.client.hub/templates/pages/redirections"
	"github.com/labstack/echo/v4"
)

func (eh echoHandlers) HomeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	sec := c.(*context.ServerEchoContext)
	if sec == nil {
		return c.String(http.StatusInternalServerError, "error getting server context")
	}

	platformClient := sec.GetPlatformClientService()
	platform, err := platformClient.GetPlatform(ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	lcs := sec.GetLinkClientService()

	link, err := lcs.GetLink(ctx, "home")
	if err != nil || link == nil {
		return extensions.Render(c, http.StatusOK, pages.Home(*platform.LayoutDescription))
	}

	if link.Status == status.StatusInactive {
		return echo.NewHTTPError(http.StatusNotFound, "Link is not active")
	}

	return extensions.Render(c, http.StatusOK, redirection_pages.Custom(*platform.LayoutDescription, *link))
}
