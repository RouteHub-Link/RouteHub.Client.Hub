package mq

import "strings"

type MQTopic string

var (
	MQE_LINK_SET     MQTopic = "link/set"
	MQE_LINK_DEL     MQTopic = "link/del"
	MQE_LINK_GET     MQTopic = "link/get"
	MQE_LINK_ALL     MQTopic = "link/all"
	MQE_PLATFORM_SET MQTopic = "platform/set"
	MQE_PLATFORM_GET MQTopic = "platform/get"
)

func (e MQTopic) AsTopic() string {
	return strings.Join([]string{"topic", string(e)}, "/")
}

func (e MQTopic) IsValid() bool {
	switch e {
	case MQE_LINK_SET, MQE_LINK_DEL, MQE_LINK_GET, MQE_LINK_ALL, MQE_PLATFORM_SET, MQE_PLATFORM_GET:
		return true
	}
	return false
}
