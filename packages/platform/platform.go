package platform

import (
	"github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts"
)

type Platform struct {
	Name               string                     `gorm:"type:varchar(255);not null;"`
	Slug               string                     `gorm:"type:varchar(255);not null;"`
	DefaultRedirection redirection.Option         `gorm:"type:varchar(255);not null;"`
	Status             status.State               `gorm:"type:varchar(255);not null;"`
	LayoutDescription  *layouts.LayoutDescription `gorm:"foreignKey:PlatformID;"`
}
