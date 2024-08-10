package groups

import (
	"net/http"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
)

var (
	mockMeta         = components.MetaDescription{Title: "RouteHub", Description: "RouteHub is a platform that allows you to create, share, and discover routes for your favorite activities."}
	navbarItems      = []components.NavbarItem{{Text: "Home", URL: "/", Target: "_self", Icon: "home"}, {Text: "About", URL: "/about", Target: "_self", Icon: "info", Dropdown: &[]components.NavbarItem{{Text: "Contact", URL: "/contact", Target: "_self", Icon: "contact_mail"}}}}
	navbarEndButtons = []components.NavbarButton{{Text: "Login", URL: "/login", Target: "_self", ColorClass: "is-secondary"}, {Text: "Sign Up", URL: "/signup", Target: "_self", ColorClass: "is-primary"}}
	brandImg         = components.ImageDescription{SRC: "https://avatars.githubusercontent.com/u/153122518?s=250", Alt: "RouteHub", Width: "30", Height: "30"}

	mavbar            = components.NavbarDescription{BrandName: "RouteHub", BrandURL: "https://routehub.link", BrandImg: &brandImg, StartItems: &navbarItems, EndButtons: &navbarEndButtons}
	layoutDescription = layouts.LayoutDescription{MetaDescription: mockMeta, NavbarDescription: nil}
)

func MapPublicRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
}

func HomeHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, pages.Home(layoutDescription))
}
