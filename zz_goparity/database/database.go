package database

// dao
import (
	"fmt"
	"log"
	"time"

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

func (db *DBDaoStruct) InitiateDBConnection(properties string) (*mgo.Session, error) {

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
	fmt.Println(time.Now().String() + ": " + "Now connected to " + db.dbType + " instance...")

	//defer session.Close()
	//fmt.Printf("Closing mongodb connection...")

	return session, err
}

func (db *DBDaoStruct) ConnectToDB(s *mgo.Session, properties string) (*mgo.Database, error) {

	// sooon, replace with properteis file stuff
	db.dbType = "mongodb"
	db.dbServer = "http://localhost"
	db.dbPort = ":27107"
	db.dbName = "goparitydb_phase1"
	db.dbUserName = ""
	db.dbPassword = ""

	var dbc *mgo.Database
	var err error
	dbc = s.DB(db.dbName)
	if dbc.Name == "test" {
		log.Fatal("Unable to find database " + db.dbName + "...")
		s.Close()
	}
	fmt.Println(time.Now().String() + ": " + "Now connected to " + db.dbType + " database " + db.dbName + "...")

	return dbc, err
}
