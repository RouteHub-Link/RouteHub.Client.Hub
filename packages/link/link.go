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
	Title              string                      `json:"title" form:"title"`
	Subtitle           string                      `json:"subtitle" form:"subtitle"`
	ContentContainer   string                      `json:"content_container" form:"content_container"`
	RedirectionURLText string                      `json:"redirection_url_text" form:"redirection_url_text"`
	RedirectionDelay   *int                        `json:"redirection_delay" form:"redirection_delay"`
	MetaDescription    *components.MetaDescription `json:"meta_description" form:"meta_description"`
	AdditionalHead     *string                     `json:"additional_head" form:"additional_head"`
	AdditionalFooter   *string                     `json:"additional_foot" form:"additional_foot"`
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
