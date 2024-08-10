package server

import (
	"github.com/RouteHub-Link/routehub.client.hub/server/router"
	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()

	router.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
