// server
//https://thenewstack.io/make-a-restful-json-api-go/

// import packages and libraries
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webstuff/restapi/v2/middleware"
	//dependencies
)

func main() {

	// set the router
	//router := mux.NewRouter().StrictSlash(true)
	// set the endpoints for rest apis
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/todos", ToDoIndex)
	// router.HandleFunc("/todos/{todoID}", ToDoShow) //variable todoid
	router := middleware.NewRouter()

	err := http.ListenAndServe(":8082", router)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Println(time.Now())
		fmt.Println("Now Serving...")
	}

}
