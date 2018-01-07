package routing

import (
	// custom libraries
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {

	// create new gorilla mux router
	router := mux.NewRouter().StrictSlash(true)

	return router
}

func SetRoutes()
