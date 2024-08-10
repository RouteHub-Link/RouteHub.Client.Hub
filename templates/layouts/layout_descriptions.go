package layouts

import (
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/a-h/templ"
)

type LayoutDescription struct {
	MetaDescription   components.MetaDescription
	NavbarDescription *components.NavbarDescription
	FooterDescription *components.FooterDescription
}

type MainDescription struct {
	AdditionalHead   *templ.Component
	MainContent      templ.Component
	AdditionalFooter *templ.Component

	LayoutDescription LayoutDescription
}

func (md MainDescription) SetHead(head templ.Component) MainDescription {
	md.AdditionalHead = &head
	return md
}

func (md MainDescription) SetFooter(footer templ.Component) MainDescription {
	md.AdditionalFooter = &footer
	return md
}

func (md MainDescription) SetLayoutDescription(ld LayoutDescription) MainDescription {
	md.LayoutDescription = ld
	return md
}

func (md MainDescription) SetMainContent(content templ.Component) MainDescription {
	md.MainContent = content
	return md
}

func (md MainDescription) GetAdditionalFooter() templ.Component {
	return *md.AdditionalFooter
}

func (md MainDescription) GetAdditionalHead() templ.Component {
	return *md.AdditionalHead
}
