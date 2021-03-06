package database

import (
	"database/sql"
	"log"
	"zz_wickedsick/utils/errorstuff"

	_ "github.com/go-sql-driver/mysql"
)

// save the dbc variable as global (reference pointer)
// needed for handlers to access the db connection
var DBC *sql.DB

// ConnectToDB connects to the backend database connected to this app
func ConnectToDB(dbConnectCommand string) {

	// variables
	var err error

	log.Printf("Now Connecting to database...")
	// connect to database
	DBC, err = sql.Open("mysql", dbConnectCommand)
	errorstuff.HandleErr(err)

	// ping to make sure database is up and running
	err = DBC.Ping()
	errorstuff.HandleErr(err)

	log.Println("Database up and running...")
}

func SelectFromDB() {
	rows, err := DBC.Query("select * from user")
	errorstuff.HandleErr(err)
	log.Println(rows)
}
