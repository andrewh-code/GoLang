package middleware

import (
	"net/http"
	"webstuff/restapi/v2/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//versioning variable
var version = "v2"

// create a new file just to handle the routes themselves
// this file focuses on the paths while the router.go file focuses on driving the routes
var routes = Routes{
	Route{
		"Index",
		"Get",
		"/api/{version}/",
		controller.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/api/{version}/todos",
		controller.ToDoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/api/{version}/todos/{todoID}",
		controller.ToDoShow,
	},
	Route{
		"TodoCreate",
		"GET", //try to reach page but get 404, because it's post? - correct ()chagne back to post
		"/api/{version}/todocreate",
		controller.TodoCreate,
	},
	Route{
		"TodoDelete",
		"GET",
		"/api/{version}/tododelete/{todoID}",
		controller.TodoDelete,
	},
}
