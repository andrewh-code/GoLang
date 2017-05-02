package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username  string `json: "username"`
	Password  string `json: "password"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

func funcUser(w http.ResponseWriter, r *http.Request) {

	log.Printf("http request: ", r)
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

func main() {

	http.HandleFunc("/", funcUser)
	fmt.Println("Now Serving...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
