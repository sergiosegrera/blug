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

// HTML responds with a static html page
func HTML(w http.ResponseWriter, r *http.Request, code int, file string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	http.ServeFile(w, r, "./front/templates/"+file)
}
