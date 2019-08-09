package web

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
	"go-delic-products/elastic"
	"go-delic-products/model"
	"io/ioutil"
	"log"
	"net/http"
)

func NewElasticWebHandler(esClient *elasticsearch.Client) *httpElastic {
	return &httpElastic{
		client:   &elastic.PostElastic{},
		esClient: esClient,
	}
}

type httpElastic struct {
	client   *elastic.PostElastic
	esClient *elasticsearch.Client
}

func (esHandler *httpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)

	postApi := esHandler.client

	idSaved, _ := postApi.Save(&post, *esHandler.esClient)

	res, _ := json.Marshal(&idSaved)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (esHandler *httpElastic) GetByIdHandler(writer http.ResponseWriter, request *http.Request) {
	id, client := mux.Vars(request)["id"], esHandler.client
	res, err := client.FindById(id, *esHandler.esClient)

	mapResponse := make(map[string]interface{})
	if err != nil {
		log.Fatal("errors in the response", err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &mapResponse)

	writer.WriteHeader(http.StatusOK)

	jsonResponse, _ := json.Marshal(mapResponse)
	_, _ = writer.Write(jsonResponse)

}
