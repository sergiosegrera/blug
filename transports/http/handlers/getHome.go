package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	HTML(w, r, 200, "post.html")
}
