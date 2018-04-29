package database

import (
	"testing"
)

func TestConnectToDB(t *testing.T) {

	var db DBDaoStruct

	result, err := db.ConnectToDB()

	if err != nil {
		t.Errorf("Error...")
	}
	result.Close()

}
