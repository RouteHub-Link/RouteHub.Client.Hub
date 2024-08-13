package server

import (
	server_context "github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/router"
	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()
	server_context.ApplyMiddleware(e)

	router.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
