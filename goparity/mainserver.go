package main

import (
	"fmt"
	"net/http"

	"goparity/app/routing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to goparity server")

}

func main() {
	http.HandleFunc("/", handler)

	port := "9000"
	// initialize router
	router := routing.InitializeRouter()

	fmt.Printf("now servering on localhost: %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
