package main

import (
	"fmt"
	"net/http"
	"zz_goparity/routing"

	"github.com/globalsign/mgo"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to goparity server")

}

var DBC *mgo.Database

func main() {
	http.HandleFunc("/", handler)

	port := ":9000"
	// initialize router
	router := &routing.Router{}
	router.InitializeRouter()
	router.SetRoutes()

	// initialize database
	// db := &database.DBDaoStruct{}
	// dbSession, err := db.InitiateDBConnection("mock properties")
	// if err != nil {

	// }
	// DBC, err = db.ConnectToDB(dbSession, "mock properties")
	// if err != nil {

	// }

	// err := http.ListenAndServe(":"+port, router)
	// if err != nil {
	// 	panic(err)
	// }
	router.Run(port)
	fmt.Printf("Now servering...")

}
