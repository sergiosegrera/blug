package models

// Post is the blug post data model
type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	HTML     string `json:"html"`
	Markdown string `json:"markdown"`
}
