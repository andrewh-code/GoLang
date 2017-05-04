package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/debug"

	"github.com/gorilla/securecookie"
)

type LoginJSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,string"`
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

var cookieName = "wickedsickcookie" //cookies can't have spaces in their names

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
	debug.FormatRequest(r)
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
		SetSession(user.UserName, w)
		redirectTarget = "/inside"
	}
	debug.Log("logincontroller.go: ", "redirecting to "+redirectTarget)
	http.Redirect(w, r, redirectTarget, 302)
}

func RedirectFailedLogin(w http.ResponseWriter, r *http.Request) {
	debug.FormatRequest(r)
	fmt.Fprintf(w, "Failed login attempt")
}

func RedirectInside(w http.ResponseWriter, r *http.Request) {
	debug.FormatRequest(r)
	userName := GetUserName(r)
	debug.Log("logincontroller.go: ", "username is "+userName)

	if userName != "" { //do other types of validation
		// load page
		t, _ := template.ParseFiles("app/view/inside.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/redirect", 302)
	}

}

func SetSession(userName string, w http.ResponseWriter) {

	value := map[string]string{
		"name": userName,
	}

	//encode the session
	if encoded, err := cookieHandler.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
	}
	debug.Log("\tlogincontroller.go-->SetSession: ", value["name"]+" "+cookieName)
}

func GetUserName(r *http.Request) (userName string) {

	if cookie, err := r.Cookie(cookieName); err == nil {
		cookieValue := make(map[string]string)

		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	debug.Log("logincontroller.go --> GetUserName ", cookieName)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

// deletes current session
func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1, //sets when the cookie expires (-1 means destroy the cookie when close the session)
	}
	http.SetCookie(w, cookie)
}
