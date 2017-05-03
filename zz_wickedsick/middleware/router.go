package middleware

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

	// // API subrouter
	// // Serves all JSON REST handlers prefixed with /api
	// s := r.PathPrefix("/api").Subrouter()
	// for _, route := range apiRoutes {
	// 	s.Methods(route.Method).
	// 		Path(route.Pattern).
	// 		Name(route.Name).
	// 		Handler(AuthorizeHandler(route.HandlerFunc))
	// }
}
