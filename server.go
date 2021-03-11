package main

import (
	"GO/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

//type Recorder interface{
//	Bind(string, User) Recorder
//	Interface() interface{}  // Get the struct that has been linked
//	Insert() error // INSERT just one record
//	Update() error // UPDATE just one record
//	Delete() error // DELETE just one record
//	Exists() (bool, error) // Check for just one record
//	ExistsWhere(cond interface{}, args ...interface{}) (bool, error)
//	Load() error  // SELECT just one record
//	LoadWhere(cond interface{}, args ...interface{}) error // Alternate Load()
//}


func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("FROM : " + c.Request().Host + "		-->	 URL : " + c.Request().RequestURI)
		if c.Request().Header.Get("Authorization") != "" {
			next(c)
			return nil
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization :(")
		//return http.Error()
	}
}

func main() {

	if _, err:=os.Stat("./database/Authen.db"); err == nil{
		log.Println("	:	Database exits.")
		db, err := sql.Open("sqlite3", "./database/Authen.db")
		if err != nil{
			log.Fatal(err.Error())
		}
		_, err = db.Exec("create table if not EXISTS user_info(username text NOT NULL PRIMARY KEY, password text NOT NULL);")
		if err != nil{
			log.Fatal(err.Error())
		}

		

	}else if os.IsNotExist(err) {
		log.Println("database not found")
		file, err := os.Create("./database/Authen.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("Authen.db created")
		db, err := sql.Open("sqlite3", "./database/Authen.db")
		_, err = db.Exec("create table if not EXISTS user_info(username text NOT NULL PRIMARY KEY, password text NOT NULL);")
		if err != nil{
			log.Fatal(err.Error())
		}
	} else {
		log.Println("ERROR")
	}

	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello, MyMiddleware)
	// e. POST("/authen/create_user", create, MyMiddleware)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	temp := models.User{1,"test","123456"}
	log.Println(temp)
	return c.String(http.StatusOK, "Hello, World!")
}
