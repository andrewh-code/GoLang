package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// refactor routing statements into a new file just to handle routes
// also makes it easier to add in a logger

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
