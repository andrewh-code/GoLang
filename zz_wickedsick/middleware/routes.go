package middleware

import (
	"net/http"
	"zz_wickedsick/app/controller"
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
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"Login",
		"GET",
		"/login",
		controller.Login,
	},
}
