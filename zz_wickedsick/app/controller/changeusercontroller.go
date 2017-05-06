package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/cookies"

	"github.com/gorilla/mux"
)

func UserInformationGETFromURI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var u model.User
	cookieUserName := cookies.GetUserName(r)
	uriUserName := vars["user"]

	log.Println("changeusercontroller.go", "UserInformationGETFromURI")
	log.Println("uriUserName is: " + uriUserName)
	log.Println("cookieUserName is: " + cookieUserName)

	if cookieUserName == uriUserName {
		u = model.GetUserDetails(uriUserName)

		fmt.Fprintln(w, u.UserName)
		fmt.Fprintln(w, u.FirstName)
		fmt.Fprintln(w, u.LastName)
		fmt.Fprintln(w, u.Email)
		fmt.Fprintln(w, u.Address)
		fmt.Fprintln(w, u.PostalCode)
		fmt.Fprintln(w, u.PhoneNumber)

	} else {
		fmt.Fprint(w, "cannot access user information")
	}
}

func UserInformationGET(w http.ResponseWriter, r *http.Request) {

	var u model.User
	cookieUserName := cookies.GetUserName(r)

	log.Println("changeusercontroller.go", "UserInformationGET")
	log.Println("cookieUserName is: " + cookieUserName)

	if len(cookieUserName) > 1 {
		u = model.GetUserDetails(cookieUserName)

		fmt.Fprintln(w, u.UserName)
		fmt.Fprintln(w, u.FirstName)
		fmt.Fprintln(w, u.LastName)
		fmt.Fprintln(w, u.Email)
		fmt.Fprintln(w, u.Address)
		fmt.Fprintln(w, u.PostalCode)
		fmt.Fprintln(w, u.PhoneNumber)

	} else {
		fmt.Fprint(w, "cannot access user information")
	}
}
func ChangeUserGET(w http.ResponseWriter, r *http.Request) {

	// TODO: figure out why the template isn't working with variables
	changeUserFile := "app/view/change.html"
	t, _ := template.ParseFiles(changeUserFile)
	t.Execute(w, nil)

}

func ChangeUserPUT(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var updatedUser model.User
	cookieUserName := cookies.GetUserName(r)

	// if cookie DOES exist
	// find a better way for cookie validation
	if len(cookieUserName) > 1 {
		r.ParseForm()
		user.UserName = cookieUserName
		user.FirstName = r.Form.Get("firstname")
		user.LastName = r.Form.Get("lastname")
		user.Email = r.Form.Get("email")
		user.PhoneNumber = r.Form.Get("phonenumber")
		user.Address = r.Form.Get("address")
		user.PostalCode = r.Form.Get("postalcode")

		// basic validation

		// change user
		user.ChangeUser()

		// get updated user information
		updatedUser = model.GetUserDetails(cookieUserName)

		fmt.Fprintln(w, updatedUser.UserName)
		fmt.Fprintln(w, updatedUser.FirstName)
		fmt.Fprintln(w, updatedUser.LastName)
		fmt.Fprintln(w, updatedUser.Email)
		fmt.Fprintln(w, updatedUser.Address)
		fmt.Fprintln(w, updatedUser.PostalCode)
		fmt.Fprintln(w, updatedUser.PhoneNumber)
	} else {
		fmt.Fprintf(w, "Unable to change user")
	}

}

