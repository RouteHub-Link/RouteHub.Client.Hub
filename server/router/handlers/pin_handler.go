package handlers

import (
	"net/http"
	"strconv"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

func (eh echoHandlers) PinHandler(c echo.Context) error {
	// get random int between 324 to 48129

	rows := []components.TableRow{{Cells: []components.TableCell{
		{Content: "<a href='https://routehub.link' target='_blank'> Rotehub.Link Home </a>", Class: ""}, {Content: randomIntAsString(300, 3000), Class: ""}}, Class: ""},
		{Cells: []components.TableCell{
			{Content: "<a href='https://routehub.link' target='_blank'> Rotehub.Test </a>", Class: ""}, {Content: randomIntAsString(300, 3000), Class: ""}}, Class: ""},
		{Cells: []components.TableCell{
			{Content: "<a href='https://routehub.link' target='_blank'> Rotehub.Test2 </a>", Class: ""}, {Content: randomIntAsString(300, 3000), Class: ""}}, Class: ""},
	}

	headers := []components.TableHeader{{Content: "URL", Class: ""}, {Content: "Clicks", Class: ""}}

	td1 := components.TableDescription{
		Headers:              headers,
		Footers:              nil,
		FooterIsSameAsHeader: true,
		Class:                "is-bordered is-striped is-narrow is-hoverable is-fullwidth",
		Rows:                 &rows,
	}

	td2 := components.TableDescription{
		Headers:              headers,
		Footers:              nil,
		FooterIsSameAsHeader: true,
		Class:                "is-bordered is-striped is-narrow is-hoverable is-fullwidth",
		Rows:                 &rows,
	}

	table1 := components.Table(td1)
	table2 := components.Table(td2)

	pinPanels := []components.PanelDescription{
		{PanelHeading: "Top Clicks", PanelColorClass: "is-primary", PanelTable: table1},
		{PanelHeading: "Our Picks", PanelColorClass: "is-info", PanelTable: table2}}

	return extensions.Render(c, http.StatusOK, pages.Pins(eh.layoutDescription, pinPanels))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomIntAsString(min, max int) string {
	return strconv.Itoa(randomInt(min, max))
}
