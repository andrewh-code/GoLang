package main

import (
	"log"
	"net/http"
	"zz_wickedsick/middleware"
	"zz_wickedsick/utils/config"
	"zz_wickedsick/utils/database"
	"zz_wickedsick/utils/debug"
	"zz_wickedsick/utils/errors"
)

//constants
const (
	FILE_NAME = "config/config.json"
)

var MainDebug = true

/*************************************************
* main server logic to run
*
 ************************************************/
func main() {

	// variables
	var server config.ServerStruct
	var db config.DBStruct
	var err error
	var dbConnectCommand string
	// initiate the router to handle the URL/api execution
	router := middleware.NewRouter() // variable to handle the middleware portion (handle the routes and the handlers)

	// initiate debug
	debug.Init(true)
	// retrieve the configuration from the config.json file
	//TODO: put all this into one call function
	server = config.RetrieveServerConfiguration(FILE_NAME)
	db = config.RetrieveDBConfiguration(FILE_NAME)

	// connect to the database
	//TODO: refactor to pass in db struct?
	dbConnectCommand = db.Database.DBUsername + ":" +
		db.Database.DBPassword + "@" + "tcp(" +
		db.Database.Hostname +
		db.Database.DBPort + ")" + "/" +
		db.Database.DBName
	database.ConnectToDB(dbConnectCommand)

	// run the server (add conditions for https)
	log.Println("Now Serving...")

	err = http.ListenAndServe(server.Server.HTTPPort, router)
	errors.HandleErr(err)

}
