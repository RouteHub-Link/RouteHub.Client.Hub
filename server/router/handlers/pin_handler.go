package handlers

import (
	"net/http"
	"strings"

	"github.com/RouteHub-Link/routehub.client.hub/packages/cusrand"
	"github.com/RouteHub-Link/routehub.client.hub/server/context"
	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
)

func (eh echoHandlers) PinHandler(c echo.Context) error {
	pins := getMockData()
	ctx := c.Request().Context()

	sec := c.(*context.ServerEchoContext)
	platformClient := sec.GetPlatformClientService()
	platform, err := platformClient.GetPlatform(ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return extensions.Render(c, http.StatusOK, pages.Pins(*platform.LayoutDescription, pins))
}

func getMockData() []components.PanelDescription {
	rows := randomRows(15)
	rows2 := randomRows(12)
	headers := headers()

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
		Rows:                 &rows2,
	}

	table1 := components.Table(td1)
	table2 := components.Table(td2)

	return []components.PanelDescription{
		{PanelHeading: "Top Clicks", PanelColorClass: "is-primary", PanelTable: table1},
		{PanelHeading: "Our Picks", PanelColorClass: "is-info", PanelTable: table2}}
}

func headers() []components.TableHeader {
	return []components.TableHeader{{Content: "URL", Class: ""}, {Content: "Clicks", Class: ""}}
}

func randomRows(count int) []components.TableRow {
	rows := make([]components.TableRow, count)
	for i := 0; i < count; i++ {
		rows[i] = components.TableRow{Cells: randomTableCells(), Class: ""}
	}
	return rows
}

func randomTableCells() []components.TableCell {
	randomString := cusrand.UniqueRandomString(10)
	count := cusrand.RandomIntAsString(300, 3000)
	link := strings.Join([]string{"<a href='https://", randomString, ".link' target='_blank'> Rotehub.Link ", randomString, "</a>"}, "")

	return []components.TableCell{
		{Content: link, Class: ""}, {Content: count, Class: ""}}
}
