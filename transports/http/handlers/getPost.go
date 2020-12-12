package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/sergiosegrera/blug/service"
)

// MakeGetPostHandler creates a http handler for the GetPost endoint of the service
func MakeGetPostHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stringID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(stringID)
		if err != nil {
			JSON(w, 400, message{"error": "Invalid id"})
			return
		}

		post, err := svc.GetPost(r.Context(), id)
		// TODO: post not exist error, 404?
		if err != nil {
			JSON(w, 404, message{"error": "Post not found"})
			return
		}

		JSON(w, 200, post)
	}
}
