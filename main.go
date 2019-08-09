package main

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
	"go-delic-products/web"
	"log"
	"net/http"
)

const port = ":8080"

var client, _ = elasticsearch.NewDefaultClient()

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/posts", savePost).Methods(http.MethodPost)
	router.HandleFunc("/api/posts/{id}", findById).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(port, router))
}

func findById(writer http.ResponseWriter, request *http.Request) {
	elasticClient := web.NewElasticWebHandler(client)
	elasticClient.GetByIdHandler(writer, request)

}

func savePost(writer http.ResponseWriter, request *http.Request) {

	elasticClient := web.NewElasticWebHandler(client)

	elasticClient.SaveHandler(writer, request)
}
