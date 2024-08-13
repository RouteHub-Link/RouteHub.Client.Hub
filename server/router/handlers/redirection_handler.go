package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	redirection_pages "github.com/RouteHub-Link/routehub.client.hub/templates/pages/redirections"
	"github.com/labstack/echo/v4"
)

const (
	RedirectionChoiceTimed   = "timed"
	RedirectionChoiceDirect  = "direct"
	RedirectionChoiceConfirm = "confirm"
	RedirectionChoiceCustom  = "custom"
	RedirectionChoiceDefault = RedirectionChoiceTimed
)

func (eh echoHandlers) HandleShortenURL(c echo.Context) error {
	ctx := c.Request().Context()

	key := c.Param("key")

	cc := c.(*context.CustomContext)
	logger := cc.GetLogger()

	logger.Log(ctx, slog.LevelDebug, "Handling shorten URL request", slog.String("key", key))

	//choice := RedirectionChoiceTimed
	//choice := RedirectionChoiceConfirm
	choice := RedirectionChoiceCustom

	logger.Log(ctx, slog.LevelDebug, "Handling shorten URL request", slog.String("Redirection Choice", choice))

	// TODO : Get the choice from the clients
	// TODO : Override eh.layoutDescription with the new layout description
	redirectionURL := "https://www.google.com"

	testTimedDesc := redirection_pages.TimedDescription{
		Title:              "We are redirecting you to Google...",
		Description:        "You will be redirected to Google in 5 seconds.",
		RedirectionDetails: "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURL:     redirectionURL,
		RedirectionURLText: "www.google.com",
		RedirectionDelay:   strconv.Itoa(5),
	}

	confirmDesc := redirection_pages.ConfirmDescription{
		PageTitle:          "Confirm Redirect",
		Title:              "Are you sure you want to redirect to Google?",
		Description:        "You will be redirected to Google.",
		RedirectionDetails: "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURL:     redirectionURL,
		RedirectionURLText: "www.google.com",
	}

	customDesc := redirection_pages.CustomDescription{
		HTML: "<h1>Custom HTML</h1>",
	}

	switch choice {
	case RedirectionChoiceTimed:
		return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, &testTimedDesc))
	case RedirectionChoiceConfirm:
		return extensions.Render(c, http.StatusOK, redirection_pages.Confirm(eh.layoutDescription, &confirmDesc))
	case RedirectionChoiceDirect:
		return HandleDirectRendering(c, redirectionURL)
	case RedirectionChoiceCustom:
		return extensions.Render(c, http.StatusOK, redirection_pages.Custom(eh.layoutDescription, &customDesc))
	default:
		return c.String(http.StatusBadRequest, "Invalid choice")
	}
}
