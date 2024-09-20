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
	mqttClientContainer     *MQTTClientContainer
	onceMQTTClientContainer sync.Once
)

type MQTTClientContainer struct {
	LinkClientService     *link.LinkClientService
	PlatformClientService *platform.PlatformClientService
}

func NewMQTTClientContainer(rc *redis.Client, logger *slog.Logger, config services.DetailsConfig) *MQTTClientContainer {
	onceMQTTClientContainer.Do(func() {
		var pcs *platform.PlatformClientService
		pcs = platform.NewPlatformClientService(rc, logger, config.PlatformId, config.SEED)

		var lcs *link.LinkClientService
		if config.SEED {
			lcs = link.NewLinkClientServiceWithSeed(rc, logger)
		} else {
			lcs = link.NewLinkClientService(rc, logger)
		}

		mqttClientContainer = &MQTTClientContainer{
			LinkClientService:     lcs,
			PlatformClientService: pcs,
		}
	})

	return mqttClientContainer
}

func GetClientContainer() (*MQTTClientContainer, error) {
	if mqttClientContainer == nil {
		return nil, errors.New("client container is nil")
	}

	return mqttClientContainer, nil
}
