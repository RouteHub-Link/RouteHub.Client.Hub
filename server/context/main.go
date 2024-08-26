package context

import (
	"log/slog"

	"github.com/RouteHub-Link/routehub.client.hub/packages"
	"github.com/RouteHub-Link/routehub.client.hub/packages/link"
	"github.com/RouteHub-Link/routehub.client.hub/packages/platform"
	"github.com/RouteHub-Link/routehub.client.hub/services"
	"github.com/labstack/echo/v4"
)

type ServerEchoContext struct {
	echo.Context
}

const (
	loggerKey         = "logger"
	clientServicesKey = "clientServices"
)

func ApplyMiddleware(e *echo.Echo) {
	lc := services.NewLoggerConfigurer(slog.LevelInfo)
	implementEchoLogger(e, lc, loggerKey)

	clientsContainer, err := packages.GetClientContainer()

	if err == nil {
		implementClientsContainer(e, clientsContainer, clientServicesKey)
	} else {
		lc.Logger.Error("Error getting client container")
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServerEchoContext{c}
			return next(cc)
		}
	})
}

func (cc *ServerEchoContext) GetLogger() *slog.Logger {
	return cc.Get(loggerKey).(*slog.Logger)
}

func (cc *ServerEchoContext) getClientContainer() *packages.ClientContainer {
	return cc.Get(clientServicesKey).(*packages.ClientContainer)
}

func (cc *ServerEchoContext) GetLinkClientService() *link.LinkClientService {
	return cc.getClientContainer().LinkClientService
}

func (cc *ServerEchoContext) GetPlatformClientService() *platform.PlatformClientService {
	return cc.getClientContainer().PlatformClientService
}
