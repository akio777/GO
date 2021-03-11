package Database

import (
	"GO/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDatabase() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("./database/Authen.db"), &gorm.Config{})
  	if err != nil {
    	log.Println("ERROR FROM INNIT DB = 17 : "+err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}