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
// this is bad design isnt' it
var routes = Routes{
	// MVC routes
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"JSONReponse",
		"GET",
		"/json",
		Response,
	},
	Route{
		"Login",
		"GET",
		"/login",
		controller.LoginGET,
	},
	Route{
		"Login",
		"POST",
		"/login",
		controller.LoginPOST,
	},
	Route{
		"Logout",
		"POST",
		"/logout",
		controller.Logout,
	},
	Route{
		"FailedLogin",
		"GET",
		"/redirect",
		controller.RedirectFailedLogin,
	},
	Route{
		"Inside",
		"GET",
		"/inside",
		controller.RedirectInside,
	},
	Route{
		"UserInformation",
		"GET",
		"/inside/info",
		controller.UserInformationGET,
	},
	// Route{
	// 	"UserInformationFromURI",
	// 	"GET",
	// 	"/inside/{user}",
	// 	controller.UserInformationGETFromURI,
	// },
	Route{
		"ChangeUserInfoGET",
		"GET",
		"/inside/changeinfo",
		controller.ChangeUserGET,
	},
	Route{
		"ChangeUserInfoPUT",
		"PUT",
		"/inside/changeinfo",
		controller.ChangeUserPUT,
	},

	Route{
		"Register",
		"GET",
		"/accounts/register",
		controller.RegisterUserGET,
	},
	Route{
		"Register",
		"POST",
		"/accounts/register",
		controller.RegisterUserPOST,
	},

	//"Secret" api routes
	Route{
		"UserDetails",
		"GET",
		"/api/{version}/{user}",
		UserDetails,
	},
}
