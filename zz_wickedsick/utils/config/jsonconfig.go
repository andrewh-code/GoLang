package config

import (
	"encoding/json"
	"os"

	"zz_wickedsick/utils/errors"
)

// initiate structs for the .json file

type ServerInfoStruct struct {
	HostName string
	UseHTTP  bool
	HTTPPort string
}

type ServerStruct struct {
	Server ServerInfoStruct
}

type DBInfoStruct struct {
	DBType     string
	DBUsername string
	DBPassword string
	DBName     string
	Hostname   string
	DBPort     string
}

type DBStruct struct {
	Database DBInfoStruct
}

// RetrieveServerConfiguration retrieves the config.json file and retrieves the server information
func RetrieveServerConfiguration(fileName string) (server ServerStruct) {

	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&server)
	errors.HandleErr(err)
	file.Close()

	// return value is server (declared at function declaration)
	return
}

// RetrieveDBConfiguration retrieves the config.json file and retrieves the DB specific information
func RetrieveDBConfiguration(fileName string) (db DBStruct) {

	// can combine this into ONE function
	// reduce the amount of open/close file into one action instead of two separate actions
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&db)
	errors.HandleErr(err)
	file.Close()

	return
}
