package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/debug"
)

type LoginJSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,string"`
}

// LoginGet serves the http login file for the user to submit their username nad password
func LoginGET(w http.ResponseWriter, r *http.Request) {
	//serve the html file
	loginFile := "app/view/login.html"
	t, _ := template.ParseFiles(loginFile)
	t.Execute(w, nil)

}

// LoginPost
func LoginPOST(w http.ResponseWriter, r *http.Request) {

	// IF the client has a JSON input then you can use a json unmarshal of the r.Body
	// BUUUTTTTT no it doesn't

	// variables
	var user model.User

	redirectTarget := "/inside"
	// parse the data
	r.ParseForm()

	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	debug.Log("loginconttroller.go: ", user.UserName)
	// check if user exists in the database
	if user.ValidateLogin() == false {
		// user does not exist
		redirectTarget = "/redirect"
	}

	//TODO: get rid of this quick hack
	if user.UserName == "admin" || user.UserName == "test" {
		redirectTarget = "/inside"
	}

	http.Redirect(w, r, redirectTarget, 302)
}

func RedirectFailedLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Failed login attempt")
}
