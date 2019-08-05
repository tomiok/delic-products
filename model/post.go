package model

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Posted  string `json:"posted"`
	Content string `json:"content"`
}
