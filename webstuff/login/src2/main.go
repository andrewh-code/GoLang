// creation of login/authentication without using gorilla/securecookies
// using the http.Cookies golang library
package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"sync"

	"github.com/satori/go.uuid"
)

var sessionStore map[string]Client //hashmap containing clietns with cookie values being keys
var storageMutex sync.RWMutex

// client struct to save if the target client session authorized
type Client struct {
	loggedIn bool
}

var loginPage = "../index.html"

func main() {

	// variables
	//loginPage := "../index.html"
	sessionStore = make(map[string]Client) //mutex for concurrent map access
	// set the routes
	http.Handle("/hello", helloWorldHandler{})
	http.Handle("/secureHello", authenticate(helloWorldHandler{}))
	http.Handle("/login", handleLogin)

	// set the server
	http.ListenAndServe(":8001", nil)
}

type helloWorldHandler struct {
	//empty struct
}

type authenticateMiddleware struct {
	wrappedHandler http.Handler
}

// struct on function
// function associated with the struct
func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello worlddddddd!!!")
}

// create handler which will supply authorization
// if authorized, user goes to the underlying handler
func (h authenticateMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if the cookie is present
	// if it isn't, continue and create a new one
	var present bool
	var client Client

	// get a cookie named "session"
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// if the cookie exists, add it to the map of cookies
	if cookie != nil {
		storageMutex.RLock()
		client, present = sessionStore[cookie.Value]
		storageMutex.RUnlock()
	} else {
		present = false
	}

	// if cookie doesn't exist, create a brand new cookie
	if present == false {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(), // generate new random encoded string- unique identifier
		}
	}
	client = Client{false}
	storageMutex.Lock()
	sessionStore[cookie.Value] = client //store new session of the cookie ["unique id", false/true]
	// client logs in for the first time, they get a cookie, the loggin value is set to false
	storageMutex.Unlock()

	// if the client isn't logged in, send him to the login page. If he is logged in, send them to the internal page
	http.SetCookie(w, cookie)
	if client.loggedIn == false { // remember, client is a custom struct
		fmt.Fprint(w, loginPage)
		// should be hosting?
		return
	}
	if client.loggedIn == true {
		h.wrappedHandler.ServeHTTP(w, r)
		return
	}
}

func authenticate(h http.Handler) authenticateMiddleware {
	return authenticateMiddleware{h}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err != http.ErrNoCookie {
			fmt.Fprint(w, err)
			return
		} else {
			err = nil
		}
	}
	var present bool
	var client Client
	if cookie != nil {
		storageMutex.RLock()
		client, present = sessionStore[cookie.Value]
		storageMutex.RUnlock()
	} else {
		present = false
	}

	if present == false {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		client = Client{false}
		storageMutex.Lock()
		sessionStore[cookie.Value] = client
		storageMutex.Unlock()
	}
	http.SetCookie(w, cookie)

	//http.SetCookie(w, cookie)
	err = r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if subtle.ConstantTimeCompare([]byte(r.FormValue("password")),
		[]byte("password123")) == 1 {
		//login user
		client.loggedIn = true
		fmt.Fprintln(w, "Thank you for logging in.")
		storageMutex.Lock()
		sessionStore[cookie.Value] = client
		storageMutex.Unlock()
	} else {
		fmt.Fprintln(w, "Wrong password.")
	}

}
