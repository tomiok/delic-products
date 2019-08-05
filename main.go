package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = "8080"

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/posts", savePost)
	log.Fatal(http.ListenAndServe(port, router))
}

func savePost(writer http.ResponseWriter, request *http.Request) {

}
