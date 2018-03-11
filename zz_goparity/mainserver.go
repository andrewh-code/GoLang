package main

import (
	"fmt"
	"net/http"
	"zz_goparity/routing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to goparity server")

}

func main() {
	http.HandleFunc("/", handler)

	port := ":9000"
	// initialize router
	router := &routing.Router{}
	router.InitializeRouter()
	router.SetRoutes()

	fmt.Printf("now servering on localhost: %s\n", port)
	// err := http.ListenAndServe(":"+port, router)
	// if err != nil {
	// 	panic(err)
	// }
	router.Run(port)
}
