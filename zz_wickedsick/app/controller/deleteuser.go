package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/cookies"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// trying r.Method case approach for this one
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		// load delete user page
		// TODO: figure out why the template isn't working with variables
		deleteUserFile := "app/view/delete.html"
		t, _ := template.ParseFiles(deleteUserFile)
		t.Execute(w, nil)
	case "POST":

		cookieUserName := cookies.GetUserName(r)
		var u model.User

		if cookieUserName != "" {
			u.UserName = cookieUserName
			u.DeleteUser()
			// wipe the cookie session
			//fmt.Fprintln(w, "successfully deleted "+u.UserName)
			Logout(w, r)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// return method in json format

	}

}
