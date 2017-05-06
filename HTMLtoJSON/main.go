package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//curl -H "Content-Type: application/json" -X POST -d '{"username":"andrew","password":"password", "firstname": "andrew", "lastname": "andrewlastname"}' http://localhost:8080/

type User struct {
	Username  string `json: "username"`
	Password  string `json: "password"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

func funcUser(w http.ResponseWriter, r *http.Request) {

	//log.Printf("http request: ", r)
	FormatRequest(r)

	// load page
	t, _ := template.ParseFiles("register.html")
	t.Execute(w, nil)

	if r.Method == "POST" {
		// parse the data
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		log.Println("R.FORM IS: ", r.Form)
		log.Printf("decoder is: ", decoder)
		// var userStruct user
		// err := decoder.Decode(&userStruct)
		// if err != nil {
		// 	panic(err)
		// }
		// defer r.Body.Close()
		// log.Println(userStruct.username, userStruct.password, userStruct.firstname, userStruct.lastname)

		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println("body is: ", string(body))

		// var t User
		// err = json.Unmarshal(body, &t)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println("fields: ", t.Username, t.Password, t.Firstname, t.Lastname)
	}
}

func JSONREsponse(w http.ResponseWriter, r *http.Request) {

	FormatRequest(r)

	if r.Method == "GET" {
		// load page
		t, _ := template.ParseFiles("register.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		var u User
		decoder := json.NewDecoder(r.Body)

		w.Header().Set("Content-Type", "application/json")

		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		fmt.Fprintf(w, u.Username)
		fmt.Fprintf(w, u.Firstname)
		fmt.Fprintf(w, u.Lastname)
		fmt.Fprintf(w, u.Password)

		if err := json.NewEncoder(w).Encode(u); err != nil {
			panic(err)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func main() {

	http.HandleFunc("/", JSONREsponse)
	fmt.Println("Now Serving...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// only way for this to work is if the application/json
	// is an acceptable format in which case chrome is not doing so
	// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8 (in the header request)

}
