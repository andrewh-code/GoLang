package controller

import (
	"fmt"
	"html"
	"net/http"
)

// Welcome message for localhost:1234/ url
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello ", html.EscapeString(r.URL.Path))
	fmt.Fprintln(w, "Welcome to project wicked sick")

}
