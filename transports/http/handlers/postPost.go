package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sergiosegrera/blug/models"
	"github.com/sergiosegrera/blug/service"
)

// MakePostPostHandler creates a http handler for the CreatePost() endpoint of the service
func MakePostPostHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Markdown string `json:"markdown"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, message{"error": "Invalid request, missing required arguments"})
			return
		}

		post := &models.Post{
			ID:       request.ID,
			Title:    request.Title,
			Markdown: request.Markdown,
		}

		err = svc.CreatePost(r.Context(), post)
		if err != nil {
			JSON(w, 500, message{"error": "Failed creating post"})
			return
		}

		JSON(w, 200, message{"data": "success"})
	}
}
