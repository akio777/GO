package Header

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("FROM : " + c.Request().Host + "		-->	 URL : " + c.Request().RequestURI)
		if c.Request().Header.Get("Authorization") != "" {
			next(c)
			return nil
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization :(")
	}
}