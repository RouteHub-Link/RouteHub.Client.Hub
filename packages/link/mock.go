package link

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	"github.com/RouteHub-Link/routehub.client.hub/packages/cusrand"
	"github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
	"github.com/google/uuid"
)

func mockLinks(ctx context.Context, logger *slog.Logger) []Link {
	links := []Link{}

	redirectionURL := "https://www.google.com"
	timedKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	testTimedDesc := LinkContent{
		Title:              "We are redirecting you to Google...",
		Description:        "You will be redirected to Google in 5 seconds.",
		RedirectionDetails: "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURL:     redirectionURL,
		RedirectionURLText: "www.google.com",
		RedirectionDelay:   strconv.Itoa(30),
	}

	confirmKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	confirmDesc := LinkContent{
		PageTitle:          "Confirm Redirect",
		Title:              "Are you sure you want to redirect to Google?",
		Description:        "You will be redirected to Google.",
		RedirectionDetails: "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURL:     redirectionURL,
		RedirectionURLText: "www.google.com",
	}

	customKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	customDesc := LinkContent{
		HTML: "<h1>Custom HTML</h1>",
	}

	links = append(links, Link{ID: uuid.New(), Key: timedKey, Options: redirection.OptionTimed, Content: &testTimedDesc})
	links = append(links, Link{ID: uuid.New(), Key: confirmKey, Options: redirection.OptionConfirm, Content: &confirmDesc})
	links = append(links, Link{ID: uuid.New(), Key: customKey, Options: redirection.OptionCustom, Content: &customDesc})

	logger.Log(ctx, slog.LevelInfo, "Links are setted from mock", slog.String("timed Key:", timedKey), slog.String("confirm Key:", confirmKey), slog.String("custom Key:", customKey))
	return links
}
