package Routes

import (
	"GO/controllers"
	abc "GO/middlewares"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Initroute() *echo.Echo{
	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Recover())
	e.GET("/", controllers.Hello, abc.MyMiddleware)

	return e
}