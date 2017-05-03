package controller

import (
	"html/template"
	"log"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/password"
)

// Index Welcome message for localhost:1234/ url
func Index(w http.ResponseWriter, r *http.Request) {
	//serve the html file
	indexFile := "app/view/index.html"
	t, _ := template.ParseFiles(indexFile)
	t.Execute(w, nil)

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
	validUserFlag := false
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

	// encrypt the password
	log.Println("encrypted password is: ", password.EncryptPassword(user.Password))

	// create the user in the database
	// sql connection should already be open from main.go (*global db variable)
	user.AddUser()

	validUserFlag = true

	if validUserFlag == true {
		// load a successful regisration page, OR
		// TODO: load a pop up window that says registration completed?
		registerCompleteFile := "app/view/registrationcomplete.html"
		t, _ := template.ParseFiles(registerCompleteFile)
		t.Execute(w, nil)
	} else {
		//reload the registration page?
	}

}

// take the user struct as an input, output an error value?
// TODO: Work on this later
func validateUser(u model.User) bool {

	// validate the username
	validUserFlag := true

	// validate the firstname
	if len(u.FirstName) < 1 || len(u.FirstName) > 64 { //due to varchar(64) constraint on db
		validUserFlag = false
	}
	// validate the last name
	if len(u.LastName) < 1 || len(u.LastName) > 64 { //due to varchar(64) constraint on db
		validUserFlag = false
	}
	if len(u.Address) < 1 || len(u.Address) > 150 {
		validUserFlag = false
	}
	// use regular expressoin to validate email
	//use regular expression to validate postal code

	return validUserFlag
}

// how to hash a password
