package link

import (
	redirection "github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/google/uuid"
)

type Link struct {
	ID      uuid.UUID          `gorm:"type:uuid;primary_key;"`
	Key     string             `gorm:"type:varchar(255);not null;"`
	Options redirection.Option `gorm:"type:varchar(255);not null;"`
	Content *LinkContent       `gorm:"type:varchar(255);not null;"`
}

type LinkContent struct {
	Title              string
	Subtitle           string
	ContentContainer   string
	RedirectionURL     string
	RedirectionURLText string
	RedirectionDelay   string
	MetaDescription    *components.MetaDescription
	AdditionalHead     *string
	AdditionalFooter   *string
}
