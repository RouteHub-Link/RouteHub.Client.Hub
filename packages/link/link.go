package link

import (
	redirection "github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
)

type Link struct {
	ID      string             `gorm:"type:uuid;primary_key;"`
	Target  string             `gorm:"type:varchar(255);not null;"`
	Options redirection.Option `gorm:"type:varchar(255);not null;"`
}
