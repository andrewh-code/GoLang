package controller

import (
	"html/template"
	"net/http"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {
	//serve the html file
	loginFile := "app/view/login.html"
	t, _ := template.ParseFiles(loginFile)
	t.Execute(w, nil)

}
func LoginPOST(w http.ResponseWriter, r *http.Request) {

}
