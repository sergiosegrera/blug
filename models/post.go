package models

type Post struct {
	Id       int `json:"id"`
	Title    string `json:"title"`
	Html     string `json:"html"`
	Markdown string `json:"markdown"`
}
