package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/RouteHub-Link/routehub.client.hub/server/extensions"
	"github.com/RouteHub-Link/routehub.client.hub/templates/layouts/components"
	"github.com/RouteHub-Link/routehub.client.hub/templates/pages"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

func (eh echoHandlers) PinHandler(c echo.Context) error {
	pins := getMockData()

	return extensions.Render(c, http.StatusOK, pages.Pins(eh.layoutDescription, pins))
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

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomIntAsString(min, max int) string {
	return strconv.Itoa(randomInt(min, max))
}

func randomRows(count int) []components.TableRow {
	rows := make([]components.TableRow, count)
	for i := 0; i < count; i++ {
		rows[i] = components.TableRow{Cells: randomTableCells(), Class: ""}
	}
	return rows
}

func randomTableCells() []components.TableCell {
	randomString := uniqueRandomString(10)
	count := randomIntAsString(300, 3000)
	link := strings.Join([]string{"<a href='https://", randomString, ".link' target='_blank'> Rotehub.Link ", randomString, "</a>"}, "")

	return []components.TableCell{
		{Content: link, Class: ""}, {Content: count, Class: ""}}
}

func uniqueRandomString(count int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, count)
	for i := range b {
		b[i] = charset[randomInt(0, len(charset))]
	}
	return string(b)
}
