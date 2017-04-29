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

type Database struct {
	DBServerName string
	DBName       string
	DBUser       string
	DBPassword   string
}

type DatabaseSlice struct {
	Databases []Database
}

func main() {
	var s Serverslice
	//str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Unmarshal([]byte(str), &s)
	//fmt.Println(s)

	// read from file
	var db DatabaseSlice
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&db)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
	fmt.Println(s)
}
