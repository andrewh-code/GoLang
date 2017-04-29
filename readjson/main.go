package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
	User       string
}

type Serverslice struct {
	Servers []Server
}

// remember, the struct field names HAVE TO match the naems in the config.json file
type Database struct {
	DBServerName string
	DBName       string
	DBUser       string
	DBPassword   string
}

// a struct to hold an array of the initial struct (just in case if there is an array of objects in the json file)
type DatabaseSlice struct {
	Databases Database
}

func decode(fileName string) (db DatabaseSlice) {

	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&db)
	if err != nil {
		panic(err)
	}
	return
	//fmt.Println(db)

}

func main() {
	//var s Serverslice
	//str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Unmarshal([]byte(str), &s)
	//fmt.Println(s)

	// read from file
	fileName := "config.json"
	var db DatabaseSlice
	db = decode(fileName)
	fmt.Println(db)

}
