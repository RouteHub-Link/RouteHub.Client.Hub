package link

import (
	"fmt"

	redirection "github.com/RouteHub-Link/routehub.client.hub/packages/redirection"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/google/uuid"
)

type Link struct {
	ID      uuid.UUID          `gorm:"type:uuid;primary_key;"`
	Target  string             `gorm:"type:varchar(255);not null;"`
	Key     string             `gorm:"type:varchar(255);not null;"`
	Options redirection.Option `gorm:"type:varchar(255);not null;"`
	Content *LinkContent       `gorm:"type:varchar(255);not null;"`
}

type LinkContent struct {
	Title              string
	Subtitle           string
	ContentContainer   string
	RedirectionURLText string
	RedirectionDelay   *int
	MetaDescription    *components.MetaDescription
	AdditionalHead     *string
	AdditionalFooter   *string
}

func (lc *LinkContent) GetRedirectionDelay() string {
	if lc.RedirectionDelay == nil {
		return ""
	}
	return fmt.Sprint(*lc.RedirectionDelay)
}

func (lc *LinkContent) SetRedirectionDelay(delay int) {
	lc.RedirectionDelay = &delay
}
