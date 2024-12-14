package link

import (
	"context"
	"log/slog"
	"strings"

	"github.com/RouteHub-Link/routehub.client.hub/packages/cusrand"
	"github.com/RouteHub-Link/routehub.client.hub/packages/enums"
)

func mockLinks(ctx context.Context, logger *slog.Logger) []Link {
	links := []Link{}

	redirectionURL := "https://www.google.com"
	timedKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	testTimedDesc := LinkContent{
		Title:              "We are redirecting you to Google...",
		Subtitle:           "You will be redirected to Google in 5 seconds.",
		ContentContainer:   "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURLText: "www.google.com",
	}
	testTimedDesc.SetRedirectionDelay(5)

	confirmKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	confirmDesc := LinkContent{
		Title:              "Are you sure you want to redirect to Google?",
		Subtitle:           "You will be redirected to Google.",
		ContentContainer:   "Google is a search engine that allows you to search for information on the internet.",
		RedirectionURLText: "www.google.com",
	}

	customKey := strings.Join([]string{"test", cusrand.UniqueRandomString(4)}, "-")
	customDesc := LinkContent{
		ContentContainer: "<h1>Custom HTML</h1>",
	}

	links = append(links, Link{Path: timedKey, Target: redirectionURL, Options: enums.RedirectionChoiceTimed, Content: &testTimedDesc})
	links = append(links, Link{Path: confirmKey, Target: redirectionURL, Options: enums.RedirectionChoiceConfirm, Content: &confirmDesc})
	links = append(links, Link{Path: customKey, Target: redirectionURL, Options: enums.RedirectionChoiceCustom, Content: &customDesc})

	logger.Log(ctx, slog.LevelInfo, "Links are setted from mock", slog.String("timed Key:", timedKey), slog.String("confirm Key:", confirmKey), slog.String("custom Key:", customKey))
	logger.Log(ctx, slog.LevelInfo, "Timed Link", slog.String("Key:", timedKey))
	logger.Log(ctx, slog.LevelInfo, "Confirm Link", slog.String("Key:", confirmKey))
	logger.Log(ctx, slog.LevelInfo, "Custom Link", slog.String("Key:", customKey))

	return links
}
