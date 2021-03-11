package main

import (
	Database "GO/database"
	"GO/routes"
	"log"
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

func main() {
	Database.DB = Database.InitDatabase()
	if Database.DB != nil{
		log.Println("DB CONNECTED !")
	}

	// if _, err:=os.Stat("./database/Authen.db"); err == nil{
	// 	log.Println("	:	Database exits.")
	// 	db, err := sql.Open("sqlite3", "./database/Authen.db")
	// 	if err != nil{
	// 		log.Fatal(err.Error())
	// 	}
	// 	_, err = db.Exec("create table if not EXISTS user_info(username text NOT NULL PRIMARY KEY, password text NOT NULL);")
	// 	if err != nil{
	// 		log.Fatal(err.Error())
	// 	}

		

	// }else if os.IsNotExist(err) {
	// 	log.Println("database not found")
	// 	file, err := os.Create("./database/Authen.db")
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	file.Close()
	// 	log.Println("Authen.db created")
	// 	db, err := sql.Open("sqlite3", "./database/Authen.db")
	// 	_, err = db.Exec("create table if not EXISTS user_info(username text NOT NULL PRIMARY KEY, password text NOT NULL);")
	// 	if err != nil{
	// 		log.Fatal(err.Error())
	// 	}
	// } else {
	// 	log.Println("ERROR")
	// }

	e := Routes.Initroute()
	e.Logger.Fatal(e.Start(":1323"))

}

// ! controllers -> function do something
// ! models -> data struct
// ! routes -> initRoute 

