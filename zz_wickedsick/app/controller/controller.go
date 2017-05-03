package controller

import (
	"html/template"
	"log"
	"net/http"
	"zz_wickedsick/app/model"
)

// Index Welcome message for localhost:1234/ url
func Index(w http.ResponseWriter, r *http.Request) {
	//serve the html file
	indexFile := "app/view/index.html"
	t, _ := template.ParseFiles(indexFile)
	t.Execute(w, nil)

}

func LoginGET(w http.ResponseWriter, r *http.Request) {
	//serve the html file
	loginFile := "app/view/login.html"
	t, _ := template.ParseFiles(loginFile)
	t.Execute(w, nil)

}
func LoginPOST(w http.ResponseWriter, r *http.Request) {

}

func RegisterUserGET(w http.ResponseWriter, r *http.Request) {

	// serve the html registration file

	registerFile := "app/view/register.html"
	t, _ := template.ParseFiles(registerFile)
	t.Execute(w, nil)
}

// TODO: Refactor this function to separeate validation and actual insertion of new user into the database
// TODO: Refactor to add a function that takes in an html form and retrieves all the input names and compares it to the specific struct
// it is supposed to insert the names into
func RegisterUserPOST(w http.ResponseWriter, r *http.Request) {

	// validate the user input
	// make sure that the username is unique
	// make sure that the user does not exist already
	// create the user in the database
	var errorMsg []string
	registrationComplete := false
	errorMsg = append(errorMsg, "Unable to Register User due to: ")

	// parse the form to retrieve the values
	r.ParseForm()
	// TODO: do backend validation (combine with front end validation)

	var user model.User

	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	user.FirstName = r.Form.Get("firstname")
	user.LastName = r.Form.Get("lastname")
	user.Email = r.Form.Get("email")
	user.PhoneNumber = r.Form.Get("phonenumber")
	user.Address = r.Form.Get("address")
	user.PostalCode = r.Form.Get("postalcode") //valdiate (6 chars long)

	log.Println(user)

	// create the user in the database
	// sql connection should already be open from main.go (*global db variable)
	user.AddUser()

	registrationComplete = true

	if registrationComplete == true {
		// load a successful regisration page, OR
		// load a pop up window that says registration completed
		registerCompleteFile := "app/view/registrationcomplete.html"
		t, _ := template.ParseFiles(registerCompleteFile)
		t.Execute(w, nil)
	} else {
		//reload the registration page?
	}

}