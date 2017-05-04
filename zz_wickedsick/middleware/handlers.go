package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

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
