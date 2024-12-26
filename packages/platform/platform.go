package platform

import (
	"github.com/RouteHub-Link/routehub.client.hub/packages/enums"
	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts"
)

type Platform struct {
	Name               string                     `gorm:"type:varchar(255);not null;" json:"name"`
	Slug               string                     `gorm:"type:varchar(255);not null;" json:"slug"`
	DefaultRedirection enums.RedirectionChoice    `gorm:"type:varchar(255);not null;" json:"default_redirection"`
	Status             status.StatusState         `gorm:"type:varchar(255);not null;" json:"status"`
	LayoutDescription  *layouts.LayoutDescription `gorm:"foreignKey:PlatformID;" json:"hub_details"`
}
