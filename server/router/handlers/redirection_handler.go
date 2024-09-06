package handlers

import (
	"log/slog"
	"net/http"

	redirection "github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
	"github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	redirection_pages "github.com/RouteHub-Link/routehub.client.hub/templates/pages/redirections"
	"github.com/labstack/echo/v4"
)

func (eh echoHandlers) HandleShortenURL(c echo.Context) error {
	ctx := c.Request().Context()

	sec := c.(*context.ServerEchoContext)
	platformClient := sec.GetPlatformClientService()
	platform, err := platformClient.GetPlatform(ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	key := c.Param("key")

	logger := sec.GetLogger()

	logger.Log(ctx, slog.LevelDebug, "Handling shorten URL request", slog.String("key", key))
	lcs := sec.GetLinkClientService()
	link, err := lcs.GetLink(ctx, key)

	if err != nil {
		logger.Log(ctx, slog.LevelInfo, "Link not found", slog.String("key", key))
		return c.String(http.StatusNotFound, "Link not found")
	}

	switch link.Options {
	case redirection.OptionTimed:
		return extensions.Render(c, http.StatusOK, redirection_pages.Timed(*platform.LayoutDescription, *link))
	case redirection.OptionConfirm:
		return extensions.Render(c, http.StatusOK, redirection_pages.Confirm(*platform.LayoutDescription, *link))
	case redirection.OptionDirectHTTP:
		return HandleDirectRendering(c, link.Target)
	case redirection.OptionCustom:
		return extensions.Render(c, http.StatusOK, redirection_pages.Custom(*platform.LayoutDescription, *link))
	default:
		return c.String(http.StatusBadRequest, "Invalid choice")
	}
}
