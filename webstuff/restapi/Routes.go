package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// create a new file just to handle the routes themselves
// this file focuses on the paths while the router.go file focuses on driving the routes
var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		ToDoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoID}",
		ToDoShow,
	},
	Route{
		"TodoCreate",
		"GET", //try to reach page but get 404, because it's post? - correct ()chagne back to post
		"/todocreate",
		TodoCreate,
	},
}
