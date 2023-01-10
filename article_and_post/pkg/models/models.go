package models

type Article struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	IsDraft bool     `json:"is_draft"`
	Tags    []string `json:"tags"`
}
