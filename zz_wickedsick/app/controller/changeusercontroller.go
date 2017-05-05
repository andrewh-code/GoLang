package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"zz_wickedsick/app/model"
	"zz_wickedsick/utils/cookies"

	"github.com/gorilla/mux"
)

func UserInformationGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var u model.User
	cookieUserName := cookies.GetUserName(r)
	uriUserName := vars["user"]

	log.Println("uriUserName is: " + uriUserName)
	log.Println("cookieUserName is: " + cookieUserName)

	if cookieUserName == uriUserName {
		u = model.GetUserDetails(uriUserName)

		if err := json.NewEncoder(w).Encode(u); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprint(w, "cannot access user information")
	}
}
