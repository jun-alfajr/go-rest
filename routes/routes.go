package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	// e.Logger.Fatal(e.Start(":2106"))

	return e
}
