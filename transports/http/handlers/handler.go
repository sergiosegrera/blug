package handlers

import (
	"encoding/json"
	"net/http"
)

type message map[string]interface{}

// JSON basic json response
func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func HTML(w http.ResponseWriter, code int, file string) {
	
}