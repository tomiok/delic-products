package main

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
	"go-delic-products/web"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/posts", savePost)
	log.Fatal(http.ListenAndServe(port, router))
}

func savePost(writer http.ResponseWriter, request *http.Request) {

	es, _ := elasticsearch.NewDefaultClient()
	web.NewElasticWebHandler(*es)
}
