package platform

import (
	"github.com/RouteHub-Link/routehub.client.hub/packages/enums"
	"github.com/RouteHub-Link/routehub.client.hub/packages/status"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
)

func MockPlatform() Platform {

	mockMeta := components.MetaDescription{Title: "RouteHub", Description: "RouteHub is a platform that allows you to create, share, and discover routes for your favorite activities."}

	brandImg := components.ImageDescription{SRC: "https://avatars.githubusercontent.com/u/153122518?s=250", Alt: "RouteHub", Width: "30", Height: "30"}
	navbarItems := []components.NavbarItem{{Text: "Home", URL: "/", Target: "_self", Icon: "home"}, {Text: "About", URL: "/about", Target: "_self", Icon: "info", Dropdown: &[]components.NavbarItem{{Text: "Contact", URL: "/contact", Target: "_self", Icon: "contact_mail"}}}}
	navbarEndButtons := []components.NavbarButton{{Text: "Login", URL: "/login", Target: "_self", ColorClass: "is-secondary"}, {Text: "Sign Up", URL: "/signup", Target: "_self", ColorClass: "is-primary"}}

	navbar := components.NavbarDescription{BrandName: "RouteHub", BrandURL: "https://routehub.link", BrandImg: &brandImg, StartItems: &navbarItems, EndButtons: &navbarEndButtons}

	socialMediaList := []components.ASocialMedia{{Icon: "facebook", Link: "https://www.facebook.com", Target: "_blank"}, {Icon: "twitter", Link: "https://www.twitter.com", Target: "_blank"}, {Icon: "instagram", Link: "https://www.instagram.com", Target: "_blank"}, {Icon: "linkedin", Link: "https://www.linkedin.com", Target: "_blank"}}
	socialMediaContainer := components.SocialMediaContainer{SocialMediaLinks: &socialMediaList, SocialMediaPeddingClass: "pt-5", SocialMediaSizeClass: "is-medium", SocialMediaColorClass: "has-text-white"}

	footer := components.FooterDescription{ShowRouteHubBranding: true, CompanyBrandingHtml: "<strong>X Company</strong> <a href=''> X Company</a> Has Rights of this site since 1111</strong>", SocialMediaContainer: &socialMediaContainer}

	layoutDescription := layouts.LayoutDescription{MetaDescription: mockMeta, NavbarDescription: &navbar, FooterDescription: &footer}

	return Platform{
		Name:               "RouteHub",
		Slug:               "routehub",
		DefaultRedirection: enums.RedirectionChoiceTimed,
		Status:             status.StatusActive,
		LayoutDescription:  &layoutDescription,
	}
}
