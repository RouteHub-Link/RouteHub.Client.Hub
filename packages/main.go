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
		var lcs *link.LinkClientService
		if config.SEED {
			lcs = link.NewLinkClientServiceWithSeed(rc, logger)
		} else {
			lcs = link.NewLinkClientService(rc, logger)
		}

		clientContainer = &ClientContainer{
			LinkClientService:     lcs,
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
