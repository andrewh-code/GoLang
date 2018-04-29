package database

import (
	"testing"
)

func TestInitiateDBConnection(t *testing.T) {

	var db DBDaoStruct

	result, err := db.InitiateDBConnection("mock properties")

	if err != nil {
		t.Errorf("Error...")
	}
	result.Close()

}

func TestConnectToDB(t *testing.T) {

	var db DBDaoStruct

	session, err := db.InitiateDBConnection("mock properties")
	if err != nil {
		t.Errorf("Error...")
	}

	result, err := db.ConnectToDB(session, "mock properties")
	if err != nil {
		t.Errorf("Error...")
	}

	result.Logout()
	session.Close()

}
