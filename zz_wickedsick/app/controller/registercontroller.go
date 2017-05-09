package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/debug"
	"zz_wickedsick/utils/password"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,string"`
}

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

	var errorMsg []string
	var err error
	validUserFlag := true
	var user model.User
	var response JSONResponse
	errorMsg = append(errorMsg, "Unable to Register User due to: \n")

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	// parse the form to retrieve the values
	r.ParseForm()

	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	user.FirstName = r.Form.Get("firstname")
	user.LastName = r.Form.Get("lastname")
	user.Email = r.Form.Get("email")
	user.PhoneNumber = r.Form.Get("phonenumber")
	user.Address = r.Form.Get("address")
	user.PostalCode = r.Form.Get("postalcode") //valdiate (6 chars long)
	// TODO: do backend validation (combine with front end validation)

	validUserFlag, err = user.UserExists()
	if validUserFlag == false || err != nil {
		errorMsg = append(errorMsg, err.Error())
	}

	// once the validation is complete, encrypt the password
	user.Salt = password.GenerateSalt()
	user.HashedPassword = password.EncryptPassword(user.Password, user.Salt)

	debug.Log("registercontroller.go", "salt is "+user.Salt)
	debug.Log("registercontroller.go", "password is "+user.HashedPassword)

	// create the user in the database
	// sql connection should already be open from main.go (*global db variable)

	if validUserFlag == true {
		validUserFlag, err = user.AddUser()
		errorMsg = append(errorMsg, "username "+user.UserName+"already exists")
	}
	if validUserFlag == true && err == nil {
		// load a successful regisration page, OR json response
		w.WriteHeader(http.StatusOK)
		response.Success = true
		response.Data = "Successfully registered user"
		//fmt.Fprintf(w, "<h1>Successfully registered user</h1>")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Success = false
		response.Data = errorMsg
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
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
