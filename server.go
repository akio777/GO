package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("FROM : " + c.Request().Host + "		-->	 URL : " + c.Request().RequestURI)
		if c.Request().Header.Get("Authorization") != "" {
			next(c)
			return nil
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please authen")
		//return http.Error()
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method} from : ${remote_ip}, uri=${uri}, status=${status}\n",
	//}))
	//e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello, MyMiddleware)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
