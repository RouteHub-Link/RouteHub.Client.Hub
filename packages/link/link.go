package link

import (
	"fmt"

	"github.com/RouteHub-Link/routehub.client.hub/packages/enums"
	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
)

type Link struct {
	Target  string                  `gorm:"type:varchar(255);not null;" json:"target"`
	Path    string                  `gorm:"type:varchar(255);not null;" json:"path"`
	Options enums.RedirectionChoice `gorm:"type:varchar(255);not null;" json:"redirection_choice"`
	Content *LinkContent            `gorm:"type:varchar(255);not null;" json:"link_content"`
	Status  status.StatusState      `gorm:"type:varchar(255);not null;" json:"status"`
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
		return "10"
	}
	return fmt.Sprint(*lc.RedirectionDelay)
}

func (lc *LinkContent) SetRedirectionDelay(delay int) {
	lc.RedirectionDelay = &delay
}
