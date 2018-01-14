package routing

// base middleware design off of: https://github.com/mingrammer/go-todo-rest-api-example

import (
	// custom libraries
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) InitializeRouter() {

	// create new gorilla mux router
	r.Router = mux.NewRouter().StrictSlash(true)

}

func (r *Router) SetRoutes() {

	r.Router.HandleFunc("/test", PrintThis).Methods("GET")
	r.Router.HandleFunc("/", HelloWorld).Methods("GET")
}

/*
	create wrapper functions (to make things look cleaner when setting routes)
	for the HTTP methods (get, put, patch, post, etc)
*/

// Get wrapper method for HTTP GET method
func (r *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, f).Methods("GET")
}

// Post wrapper method for HTTP POST method
func (r *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, f).Methods("POST")
}

// Put wrapper method for HTTP PUT method
func (r *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wrapper method for HTTP DELETE method
func (r *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, f).Methods("DELETE")
}

func PrintThis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "print this")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func (r *Router) Run(port string) {
	err := http.ListenAndServe(port, r.Router)
	if err != nil {
		panic(err)
	}
}
