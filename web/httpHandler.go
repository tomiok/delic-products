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

	if err != nil {
		log.Fatal("errors in the response", err)
	}

	defer res.Body.Close()

	jsonResponse := make(map[string]interface{})

	body, _ := ioutil.ReadAll(res.Body)

	_ = json.Unmarshal(body, &jsonResponse)

	writer.WriteHeader(http.StatusOK)

	httpResponse, _ := json.Marshal(jsonResponse)
	_, _ = writer.Write(httpResponse)
}

func (esHandler httpElastic) FindByCriteriaHandler(writer http.ResponseWriter, request *http.Request) {
	postApi := esHandler.api

	reqBody, _ := ioutil.ReadAll(request.Body)
	m := make(map[string] interface{})
	_ = json.Unmarshal(reqBody, &m)

	query, _ := json.Marshal(m)

	log.Print(query)
	res, err := postApi.FindByCriteria(string(query), *esHandler.esClient)

	if err != nil {
		log.Fatal("errors in the response", err)
	}


	jsonResponse := make(map[string]interface{})

	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &jsonResponse)

	defer res.Body.Close()
	httpResponse, _ := json.Marshal(jsonResponse)

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(httpResponse)
}
