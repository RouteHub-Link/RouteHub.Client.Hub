package packages

import (
	"errors"
	"log/slog"
	"sync"

	"github.com/RouteHub-Link/routehub.client.hub/packages/link"
	"github.com/RouteHub-Link/routehub.client.hub/packages/platform"
	"github.com/RouteHub-Link/routehub.client.hub/services"
	"github.com/redis/go-redis/v9"
)

var (
	clientContainer     *ClientContainer
	onceClientContainer sync.Once
)

type ClientContainer struct {
	LinkClientService     *link.LinkClientService
	PlatformClientService *platform.PlatformClientService
}

func NewClientContainer(rc *redis.Client, logger *slog.Logger, config services.DetailsConfig) *ClientContainer {
	onceClientContainer.Do(func() {
		clientContainer = &ClientContainer{
			LinkClientService:     link.NewLinkClientService(rc, logger, config.SEED),
			PlatformClientService: platform.NewPlatformClientService(rc, logger, config.PlatformId, config.SEED),
		}
	})

	return clientContainer
}

func GetClientContainer() (*ClientContainer, error) {
	if clientContainer == nil {
		return nil, errors.New("client container is nil")
	}

	return clientContainer, nil
}
