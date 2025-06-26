package utils

import (
	"encoding/json"
	"net/http"

	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/models"
)

func WriteAsJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	WriteAsJson(w, struct { 
		Error string `json:"error"`
		} { Error: err.Error() })
}

func ResponseWithError(w http.ResponseWriter, status int, message string) {
	var error models.Error
	error.Message = message
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}


/*
func ResponseWithError() {
	
}

func ResponseJSON() {
	
}
*/