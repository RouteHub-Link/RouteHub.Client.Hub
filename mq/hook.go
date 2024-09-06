package mq

import (
	"bytes"
	"fmt"

	"github.com/RouteHub-Link/routehub.client.hub/mq/handlers"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/redis/go-redis/v9"
)

// Options contains configuration settings for the hook.
type MQTTHookOptions struct {
	Server      *mqtt.Server
	RedisClient *redis.Client
}

type MQTTHook struct {
	mqtt.HookBase
	config          *MQTTHookOptions
	LinkHandler     handlers.MQHandler
	PlatformHandler handlers.MQHandler
}

func (h *MQTTHook) ID() string {
	return "platform-link-hooks"
}

func (h *MQTTHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
		mqtt.OnPublished,
		mqtt.OnPublish,
	}, []byte{b})
}

func (h *MQTTHook) Init(config any) error {
	h.Log.Info("initialised")
	if _, ok := config.(*MQTTHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*MQTTHookOptions)
	if h.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}

	h.LinkHandler = handlers.NewLinkHandlers(h.config.RedisClient, h.Log)
	h.PlatformHandler = handlers.NewPlatformHandlers(h.config.RedisClient, h.Log)

	return nil
}

// subscribeCallback handles messages for subscribed topics
func (h *MQTTHook) subscribeCallback(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	h.Log.Info("hook subscribed message", "client", cl.ID, "topic", pk.TopicName)
}

func (h *MQTTHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	h.Log.Info("client connected", "client", cl.ID)

	// Example demonstrating how to subscribe to a topic within the hook.
	//h.config.Server.Subscribe("hook/direct/publish", 1, h.subscribeCallback)

	// Example demonstrating how to publish a message within the hook
	//err := h.config.Server.Publish("hook/direct/publish", []byte("packet hook message"), false, 0)
	//if err != nil {
	//	h.Log.Error("hook.publish", "error", err)
	//}

	return nil
}

func (h *MQTTHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire)
	}

}

func (h *MQTTHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	h.Log.Info(fmt.Sprintf("subscribed qos=%v", reasonCodes), "client", cl.ID, "filters", pk.Filters)
}

func (h *MQTTHook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info("unsubscribed", "client", cl.ID, "filters", pk.Filters)
}

func (h *MQTTHook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info("published to client", "client", cl.ID, "payload", string(pk.Payload))
}
