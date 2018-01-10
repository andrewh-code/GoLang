package routing

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
}

// print this function
func PrintThis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "print this")
}
