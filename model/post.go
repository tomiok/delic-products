package model

const INDEX = "shared_post"

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Posted  string `json:"posted"`
	Content string `json:"content"`
}

func (p Post) GetIndexName() string {
	return INDEX
}
