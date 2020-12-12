package handlers

import "net/http"

func MakeGetHTMLHandler(file string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		HTML(w, r, 200, file)
	}
}
