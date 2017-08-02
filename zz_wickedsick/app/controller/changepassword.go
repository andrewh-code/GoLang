package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/cookies"
	"zz_wickedsick/utils/debug"
)

func ChangePasswordGET(w http.ResponseWriter, r *http.Request) {

	// TODO: figure out why the template isn't working with variables
	changePasswordFile := "app/view/password.html"
	t, _ := template.ParseFiles(changePasswordFile)
	t.Execute(w, nil)

}

func ChangePasswordPOST(w http.ResponseWriter, r *http.Request) {

	var oldPassword string
	var newPassword string
	var validate = false
	var u model.User
	u.UserName = cookies.GetUserName(r)

	if len(u.UserName) > 1 {
		r.ParseForm()
		oldPassword = r.Form.Get("oldpassword")
		newPassword = r.Form.Get("newpassword")
	}

	validate = checkPasswords(oldPassword, newPassword)
	u.Password = oldPassword
	validate = u.ValidateLogin()

	if validate == true {
		//change password
		u.Password = newPassword
		u.ChangePassword()
		// TODO: if changepassword spits out an error, chagne validate to false
		fmt.Fprintf(w, "password successfully changed to "+u.Password)
	} else {
		fmt.Fprintf(w, "try a new password")
	}
}

func checkPasswords(oldPassword string, newPassword string) (validate bool) {
	debug.Log("changepassword.go --> checkPasswords()", "checking password validation")
	validate = false

	if (len(newPassword) > 1) &&
		(len(newPassword) <= 64) &&
		(oldPassword != newPassword) { // password not blank (do a check for password length too) and password contains numbers, variables, etc
		validate = true
	}

	return validate
}
