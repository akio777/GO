package controllers

import (
	"GO/database"
	"GO/models"
	// "fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


func Hello(c echo.Context) error {
	// temp := models.User{1,"test","123456"}
	// log.Println(temp)
	return c.String(http.StatusOK, "Hello, World!")
}

func AddUser(c echo.Context) error{
	var users []models.User
    Database.DB.Find(&users)
    log.Println("User list : ", users)
	
	// Database.DB.Create(&models.User{Id: 2,Username: "admin",Password:"123456"})
	// log.Println("KUY")
	return c.String(http.StatusOK, "CREATED ! ! !")
}


