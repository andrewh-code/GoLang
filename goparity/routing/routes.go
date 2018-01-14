package routing

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//versioning variable
var version = "v1"

// create a new file just to handle the routes themselves
// this file focuses on the paths while the router.go file focuses on driving the routes
// this is bad design isnt' it
var routes = Routes{
	// MVC routes
	Route{
		"Index",
		"GET",
		"/",
		Response,
	},
}

func Response(w http.ResponseWriter, r *http.Request) {

}
