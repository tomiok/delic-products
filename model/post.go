package model

import "time"

type Post struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Date time.Time `json:"posted"`
	Body string `json:"body"`
	
}
