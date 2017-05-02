package mydatabase

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// global variable
// declrea db connection as a pointer (need to reference it in other files)
var (
	DBC *sql.DB
)

const (
	DB_USER  = "root"
	DB_NAME  = "GODB_SANDBOX"
	PASSWORD = "password"
	DB_HOST  = "tcp(localhost:3333)"

	DB_TABLE = "user"
)

func Connect() {

	var err error

	dataSourceName := DB_USER + ":" + PASSWORD + "@" + DB_HOST + "/" + DB_NAME
	DBC, err = sql.Open("mysql", dataSourceName) // DO NOT USE :=, you're initializing a new LOCAL variable if you do this
	if err != nil {
		panic(err)
	}
	log.Println(DBC)
	//ping
	if err = DBC.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connected")
}

func Query() {
	log.Println("now querying")
	rows, err := DBC.Query("Select * from user")
	if err != nil {
		panic(err)
	}
	log.Println(rows)
}
