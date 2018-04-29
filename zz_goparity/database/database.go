package database

// dao
import (
	"log"

	"github.com/globalsign/mgo"
)

type DBDao struct {
	dbType     string
	dbServer   string
	dbPort     string
	dbName     string
	dbUserName string
	dbPassword string
}

var Connection *mgo.Database

func (db *DBDao) ConnectToDB() {

	// sooon, replace with properteis file stuff
	db.dbType = "mongodb"
	db.dbServer = "http://localhost"
	db.dbPort = ":27107"
	db.dbName = "goparitydb_phase1"
	db.dbUserName = ""
	db.dbPassword = ""

	dbURL := "mongodb://12.0.0.1:27107/"
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	Connection = session.DB(db.dbName)
}
