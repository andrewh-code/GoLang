package database

import (
	"database/sql"
	"log"
	"zz_wickedsick/utils/errors"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectToDB connects to the backend database connected to this app
func ConnectToDB(dbConnectCommand string) {

	log.Printf("Now Connecting to database...")
	// connect to database
	dbc, err := sql.Open("mysql", dbConnectCommand)
	errors.HandleErr(err)

	// ping to make sure database is up and running
	err = dbc.Ping()
	errors.HandleErr(err)

	log.Println("Database up and running...")
}
