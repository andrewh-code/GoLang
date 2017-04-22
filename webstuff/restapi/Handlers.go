package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello ", html.EscapeString(r.URL.Path))

}

func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	//todos := Todos{
	//Todo{Name: "Write presentation"},
	//Todo{Name: "Host meetup"},
	//Todo{Name: "Get tires changed"},
	//Todo{Name: "Create a restful api todo list"},
	//}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8") // tell the client to expect json encoded results
	w.WriteHeader(http.StatusOK)                                      //set the status code to 200

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ToDoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	todoIDInt, err := strconv.Atoi(todoID)

	if err != nil {
		fmt.Fprintln(w, "unable to find", todoID)
		panic(err)
	}

	for _, tStruct := range todos {
		if tStruct.Id == todoIDInt {
			fmt.Fprintln(w, tStruct.Id)
			fmt.Fprintln(w, tStruct.Name)
			fmt.Fprintln(w, tStruct.Completed)
			fmt.Fprintln(w, tStruct.Due)
		}
	}

}

//crate a new function, create to do
// go to the routes file and create the necessary info for the route, come to the handlers.go file to create the endpoint
// and figure out what to do
func TodoCreate(w http.ResponseWriter, r *http.Request) {

	var todo Todo

	// open up the body of the request (use limitreader to protect against malicious attacks) and read it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) //?
	if err != nil {
		panic(err)
	}
	// if fail, close the body of the request (think open/close file?)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	// take the data and put it into the todo struct
	// if fails, output the 422 error code (in json format)
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) //422 status code - unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated) //if the todo item was successfully created, send back status 201
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
