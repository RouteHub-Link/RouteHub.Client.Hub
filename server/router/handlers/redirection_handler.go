package handlers

import (
	"net/http"
	"strconv"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	redirection_pages "github.com/RouteHub-Link/routehub.client.hub/templates/pages/redirections"
	"github.com/labstack/echo/v4"
)

const (
	RedirectionChoiceTimed   = "timed"
	RedirectionChoiceDirect  = "direct"
	RedirectionChoiceConfirm = "confirm"
	RedirectionChoiceNot     = "not"
	RedirectionChoiceDefault = RedirectionChoiceTimed
)

func (eh echoHandlers) HandleShortenURL(c echo.Context) error {
	key := c.Param("key")

	c.Logger().Infof("Key: %s", key)

	choice := RedirectionChoiceConfirm

	c.Logger().Infof("Choice: %s", choice)

	// TODO - Get the choice from the clients
	// TODO - Override eh.layoutDescription with the new layout description

	testTimedDesc := redirection_pages.TimedDescription{
		Title:              "We are redirecting you to Google...",
		Description:        "You will be redirected to Google in 5 seconds.",
		RedirectionDetails: "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURL:     "https://www.google.com",
		RedirectionURLText: "www.google.com",
		RedirectionDelay:   strconv.Itoa(5),
	}

	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, &testTimedDesc))

	//switch choice {
	//case RedirectionChoiceTimed:
	//	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, testTimedDesc))
	//case RedirectionChoiceDirect:
	//	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, testTimedDesc))
	//case RedirectionChoiceConfirm:
	//	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, testTimedDesc))
	//case RedirectionChoiceNot:
	//	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, testTimedDesc))
	//default:
	//	return extensions.Render(c, http.StatusOK, redirection_pages.Timed(eh.layoutDescription, testTimedDesc))
	//}
}
