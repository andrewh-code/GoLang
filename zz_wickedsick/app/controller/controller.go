package controller

import (
	"html/template"
	"net/http"
)

// Index Welcome message for localhost:1234/ url
func Index(w http.ResponseWriter, r *http.Request) {

	//serve the html file
	indexFile := "app/view/index.html"
	t, _ := template.ParseFiles(indexFile)
	t.Execute(w, nil)

}

func Login(w http.ResponseWriter, r *http.Request) {

	//serve the html file
	loginFile := "app/view/login.html"
	t, _ := template.ParseFiles(loginFile)
	t.Execute(w, nil)

}
