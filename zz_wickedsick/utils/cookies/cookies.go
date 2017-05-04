package cookies

import (
	"net/http"
	"zz_wickedsick/utils/debug"

	"github.com/gorilla/securecookie"
)

// global variables
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

var cookieName = "wickedsickcookie" //cookies can't have spaces in their names

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

// deletes current session
func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1, //sets when the cookie expires (-1 means destroy the cookie when close the session)
	}
	http.SetCookie(w, cookie)
}
