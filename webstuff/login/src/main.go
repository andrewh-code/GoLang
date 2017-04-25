// https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()

func indexPageHandler(w http.ResponseWriter, r *http.Request) {

	//serve the html file
	t, _ := template.ParseFiles("../index.html")
	t.Execute(w, nil)

	fmt.Println("now serving index.html")

}

func internalPageHandler(w http.ResponseWriter, r *http.Request) {
	// extract username from request
	// if the username exists, then show the internal page. If not, redirect back to the login page
	userName := getUserName(r)
	page := "../internal.html"

	if userName != "" {
		fmt.Fprintf(w, page, userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

//PAST request handler for login
// reads the name from the POST operation. If credentials pass, username stored in session and redirect to internal sent
// else setnt to redirect page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")         // why not r.Form.get?
	password := r.FormValue("password") //difference between r.Form.get and r.FormValue?
	// you can use FormValue() or r.Form["name"]. Using FormValue(), golang calls ParseForm automatically
	// will return the first found value with the same name. If there isn't any, returns empty string

	redirectTarget := "/"
	// iff successfully pass credentials
	if name != "" && password != "" {
		// check credentials here
		setSession(name, w)
		redirectTarget = "/internal"
	}
	//fmt.Fprintf(w, "log in failed, please try again")
	http.Redirect(w, r, redirectTarget, 302)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

// func redirect(w http.ResponseWriter, r *http.Request) {
// 	//serve the html file
// 	t, _ := template.ParseFiles("../redirect.html")
// 	t.Execute(w, nil)

// 	fmt.Println("now serving redirect.html")
// }

// puts provided user name into a string map
// secure cookie handler used to encode the value map (encrypted session)
// session value stored in a cookie instance
func setSession(userName string, w http.ResponseWriter) {

	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}

}

// cookie read from the request, then the secure cookie handler is used to decrypt the cookie value
// string map and user name should be returned
func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

// deletes current session
func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func main() {
	fmt.Println("Now Serving...")
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	//router.HandleFunc("/redirect", redirect)

	http.Handle("/", router)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("unable to serve...")
		panic(err)
	}

}
