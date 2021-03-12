package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var json_map = make(map[string]interface{})

func Login(c echo.Context) error{
	log.Println()
	
		
		// return c.String(http.StatusBadRequest, "Body ERROR")
	return c.String(http.StatusOK, "Hello, World!")

}