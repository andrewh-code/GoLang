package controller

import (
	"html/template"
	"net/http"
)

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
}
