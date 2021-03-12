package Routes

import (
	"GO/controllers"
	abc "GO/middlewares"
	"github.com/labstack/echo/v4"
)

func Initroute() *echo.Echo{
	e := echo.New()
	e.HideBanner = true
	// Middleware
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	// e.Use(middleware.Recover())
	e.GET("/", controllers.Hello, abc.MyMiddleware)
	e.POST("/authen/login", controllers.Login)
	e.POST("/authen/add_user", controllers.AddUser)

	return e
}