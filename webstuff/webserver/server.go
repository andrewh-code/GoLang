package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Now serving /upload.html")
	fmt.Println("Method: ", r.Method)

	if r.Method == "GET" {
		crutime := time.Now().Unix() // mark the time when the page is retrieved
		hash := md5.New()            //help generate token (duplicate submissions prevented)

		io.WriteString(hash, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)

		fmt.Println("curtime: ", crutime)
		fmt.Println("token:", token)
		x := 32 << 20
		fmt.Println(x)
	} else {
		fmt.Println("Method: ", r.Method)

		r.ParseMultipartForm(32 << 20) // maximum memory
		x := 32 << 20
		fmt.Println(x)

		file, handler, err := r.FormFile("uploadfile") //get the file handle of the file being uploaded ("uploadfile") name attribute of file in html file
		// check if file being uploaded is a .txt file
		if string(handler.Filename[len(handler.Filename)-4:]) != ".txt" {
			panic("file uploaded is not a text file")
		}
		//check to see the type of file
		buffer := make([]byte, 512) //set 512 bytes as it takes the first 512 bytes to determine what kind of file it is
		filetype := http.DetectContentType(buffer)
		if filetype != "application/octet-stream" {
			panic("file uploaded is not of type application/octet-stream")
		}
		if err != nil {
			fmt.Println("FomrFile error:", err)
			panic(err)
		}
		defer file.Close()
		// if the file is too big for the max memory, rest of the data is saved in a temporary file
		// use r.formfile to get the file handle and use io.copy to to save the file to the system
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("tempfiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) //temporary file created
		if err != nil {
			fmt.Println("temporary file error:", err)
			panic(err)
		}
		defer f.Close()
		io.Copy(f, file)
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
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
