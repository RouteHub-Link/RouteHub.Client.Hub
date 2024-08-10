package handlers

import "github.com/RouteHub-Link/routehub.client.hub/templates/layouts"

type echoHandlers struct {
	layoutDescription layouts.LayoutDescription
}

func NewEchoHandlers(layoutDescription layouts.LayoutDescription) echoHandlers {
	return echoHandlers{layoutDescription: layoutDescription}
}
