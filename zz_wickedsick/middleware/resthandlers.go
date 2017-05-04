package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	Because we are not using a front end framework (either Ember, Angular, React, etc), I am dividing the
	[backend] REST API handlers into their own directory app/controller. They will act as the front end part
	and will also act as the back end REST API handlers
*/

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,string"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	// set the server response as JSON
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	// initialize the response struct to hold the JSON data
	var response JSONResponse

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		response.Success = true
		response.Data = "Supported Method"
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error sending reponse: %s", err)
		}
	case "POST":
		w.WriteHeader(http.StatusBadRequest)
		response.Success = false
		response.Data = "Unspoorted Method"
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error sending reponse: %s", err)
		}
	}
}

func UserDetails(w http.ResponseWriter, r *http.Request) {

}
