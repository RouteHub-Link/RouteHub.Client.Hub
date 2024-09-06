package mq

import (
	"fmt"
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *MQTTHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	h.Log.Info("received from client", "client", cl.ID, "payload", string(pk.Payload), "topic", pk.TopicName)

	pkx := pk
	if string(pk.Payload) == "hello" {
		pkx.Payload = []byte("hello world")
		h.Log.Info("received modified packet from client", "client", cl.ID, "payload", string(pkx.Payload))
	}

	mqTopic := MQTopic(strings.TrimPrefix(pk.TopicName, "topic/"))

	if !mqTopic.IsValid() {
		h.Log.Error("invalid event", "event", string(pk.Payload))
		pkx := pk
		pkx.Payload = []byte("invalid event")

		return pkx, fmt.Errorf("invalid event")
	}

	var err error

	switch mqTopic {
	case MQE_LINK_SET:
		err = h.LinkHandler.Set(pk.Payload)
	case MQE_LINK_DEL:
		err = h.LinkHandler.Delete(pk.Payload)
	case MQE_LINK_GET:
		pkx.Payload, err = h.LinkHandler.Get(pk.Payload)
	case MQE_LINK_ALL:
		pkx.Payload, err = h.LinkHandler.Fetch(pk.Payload)
	case MQE_PLATFORM_SET:
		err = h.PlatformHandler.Set(pk.Payload)
	case MQE_PLATFORM_GET:
		pkx.Payload, err = h.PlatformHandler.Get(pk.Payload)
	}

	return pkx, err
}
