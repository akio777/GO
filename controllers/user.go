package controllers

import (
	"GO/models"
	// "fmt"
	"log"
	"net/http"
	// "os"
	"github.com/labstack/echo/v4"
)


func Hello(c echo.Context) error {
	temp := models.User{1,"test","123456"}
	log.Println(temp)
	return c.String(http.StatusOK, "Hello, World!")
}

