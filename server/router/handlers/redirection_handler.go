package handlers

import (
	"log/slog"
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/packages/enums"
	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
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

	if link.Status != status.StatusActive {
		logger.Log(ctx, slog.LevelInfo, "Link is not active", slog.String("key", key))
		return c.String(http.StatusNotFound, "Link is not active")
	}

	if link.Content != nil && link.Content.MetaDescription != nil && link.Content.MetaDescription.Locale != "" {
		c.Request().Header.Set("Content-Language", link.Content.MetaDescription.Locale)
	}

	switch link.Options {
	case enums.RedirectionChoiceTimed:
		return extensions.Render(c, http.StatusOK, redirection_pages.Timed(*platform.LayoutDescription, *link))
	case enums.RedirectionChoiceConfirm:
		return extensions.Render(c, http.StatusOK, redirection_pages.Confirm(*platform.LayoutDescription, *link))
	case enums.RedirectionChoiceDirectHTTP:
		return HandleDirectRendering(c, link.Target)
	case enums.RedirectionChoiceCustom:
		return extensions.Render(c, http.StatusOK, redirection_pages.Custom(*platform.LayoutDescription, *link))
	default:
		return c.String(http.StatusBadRequest, "Invalid choice")
	}
}
