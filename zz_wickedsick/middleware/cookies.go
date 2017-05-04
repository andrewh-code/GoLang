package middleware

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

// global variables
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

var cookieName = "wicked sick cookie"

func SetSession(userName string, w http.ResponseWriter) {

	value := map[string]string{
		"name": userName,
	}

	//encode the session
	if encoded, err := cookieHandler.Encode(userName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
	}
}

func GetUserName(r *http.Request) (userName string) {

	if cookie, err := r.Cookie(cookieName); err == nil {
		cookieValue := make(map[string]string)

		if err = cookieHandler.Decode(cookieName, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return
}
