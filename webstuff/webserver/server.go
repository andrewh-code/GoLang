package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Println("\n")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Webserver Initiated!") // send data to client side
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() //retrieves everything from the page (in hashmap format?)
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("please input a username")
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func validate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now serving /validate")
	fmt.Println("method:", r.Method)

	//serve the html file
	t, _ := template.ParseFiles("validate.html")
	t.Execute(w, nil)

	//retrieve data if request is a POST
	if r.Method == "POST" {

		r.ParseForm() //always need this to get data

		fmt.Println(r.Form.Get("input1"))
		if len(r.Form.Get("input1")) == 0 {
			fmt.Println("please input a value for input1")
		}

		getint, err := strconv.Atoi(r.Form.Get("numberInput"))
		if err != nil {
			fmt.Println(err)
		}
		if (r.Form.Get("numberInput") == " ") || (r.Form.Get("numberInput") == "") {
			fmt.Println("please enter an input for numberInput")
		}

		if getint > 100 {
			fmt.Println("number is too big: ", getint)
		}

		//validate e-mail with regex
		// can validate with html5 too
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("incorrect email")
		} else {
			fmt.Println("correct email")
		}
	}
}

func dropdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now serving /dropdown.html")
	fmt.Println("Method: ", r.Method)

	//serve the html file
	t, _ := template.ParseFiles("dropdown.html")
	t.Execute(w, nil)

	if r.Method == "POST" {

		r.ParseForm() //always need this to get data

		//parse the data from the form
		slice := []string{"double chocolate", "vanilla", "strawberry", "caramel"}

		fmt.Println(r.Form.Get("Icecream Flavours"))

		//sanitize input for dropdown
		for _, selected := range slice {
			if selected == r.Form.Get("Icecream Flavours") {
				fmt.Println("ice cream flavour selected")
			} else {
				fmt.Println("ice cream flavour not selected")
			}
		}
	}

}

func checkbox(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now serving /checkbox.html")
	fmt.Println("Method: ", r.Method)

	//serve the html file
	t, _ := template.ParseFiles("checkbox.html")
	t.Execute(w, nil)
	r.ParseForm() //always need this to get data

	//validate checkbox
	sliceValidate := []string{"football", "basektball", "tennis"}
	selected := r.Form.Get("interest")

	for _, v := range sliceValidate {
		if selected == v {
			fmt.Println("selected is in the validated list")
		}
	}
}

func main() {
	t := time.Now()
	fmt.Println("Now serving...", t)
	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/login", login)   //login page
	http.HandleFunc("/validate", validate)
	http.HandleFunc("/dropdown", dropdown)
	http.HandleFunc("/checkbox", checkbox)

	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
