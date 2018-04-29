package database

// dao
import (
	"log"

	"fmt"

	"github.com/globalsign/mgo"
)

type DBDaoStruct struct {
	dbType     string
	dbServer   string
	dbPort     string
	dbName     string
	dbUserName string
	dbPassword string
}

var DBC *mgo.Database

func (db *DBDaoStruct) ConnectToDB() (*mgo.Session, error) {

	// sooon, replace with properteis file stuff
	db.dbType = "mongodb"
	db.dbServer = "http://localhost"
	db.dbPort = ":27107"
	db.dbName = "goparitydb_phase1"
	db.dbUserName = ""
	db.dbPassword = ""

	//dbURL := "mongodb://12.0.0.1:27107/"
	dbURL := "localhost:27017"
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Now connected to mongodb...")
	//defer session.Close()
	//fmt.Printf("Closing mongodb connection...")

	//DBC = session.DB(db.dbName)

	return session, err
}
