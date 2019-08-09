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

func NewElasticWebHandler(esClient *elasticsearch.Client) httpElastic {
	return httpElastic{
		api:      &elastic.PostElastic{},
		esClient: esClient,
	}
}

type httpElastic struct {
	api      *elastic.PostElastic
	esClient *elasticsearch.Client
}

func (esHandler httpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)

	postApi := esHandler.api

	idSaved, _ := postApi.Save(&post, *esHandler.esClient)

	idResponse := model.IdResponse{Id: idSaved}

	res, _ := json.Marshal(&idResponse)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (esHandler httpElastic) FindById(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	postApi := esHandler.api

	res, err := postApi.FindById(id, *esHandler.esClient)

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

func (esHandler httpElastic) FindByCriteriaHandler(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	req := make(map[string]interface{})
	_ = json.Unmarshal(requestBody, &req)

	postApi := esHandler.api



}
