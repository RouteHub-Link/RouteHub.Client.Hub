package handlers

import "github.com/labstack/echo/v4"

func HandleDirectRendering(c echo.Context, link string) error {
	return c.Redirect(302, link)
}
