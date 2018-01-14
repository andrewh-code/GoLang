package main

import (
	"encoding/json"
	"fmt"
	"json_response/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func JsonResponse(w http.ResponseWriter, r *http.Request) {
	address := model.Address{City: "Toronto", Street: "Yonge Street"}
	person := model.Person{Firstname: "Andrew", Lastname: "Code", Address: &address}

	//personInstance := person.PopulatePerson()

	json.NewEncoder(w).Encode(person)

	//fmt.Fprintf(w, "json response")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helllo world")
}
func main() {

	// set new route
	router := mux.NewRouter()

	// create handle functions
	router.HandleFunc("/json", JsonResponse).Methods("GET")
	router.HandleFunc("/", HelloWorld).Methods("GET")
	err := http.ListenAndServe(":9001", router)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Println(time.Now())
		fmt.Println("Now Serving...")
	}
}
