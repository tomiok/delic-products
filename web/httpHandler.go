package web

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
	"go-delic-products/elastic"
	"go-delic-products/model"
	"io/ioutil"
	"log"
	"net/http"
)

func NewElasticWebHandler(esClient *elasticsearch.Client) HttpElastic {
	return HttpElastic{
		api:      &elastic.PostElastic{Client: *esClient},
	}
}

type HttpElastic struct {
	api      *elastic.PostElastic
}

func (esHandler HttpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)

	postApi := esHandler.api

	idSaved, _ := postApi.Save(&post)

	idResponse := model.IdResponse{Id: idSaved}

	res, _ := json.Marshal(&idResponse)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (esHandler HttpElastic) FindById(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	postApi := esHandler.api

	res, err := postApi.FindById(id)

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

func (esHandler HttpElastic) FindByCriteriaHandler(writer http.ResponseWriter, request *http.Request) {
	postApi := esHandler.api

	res, err := postApi.FindByCriteria(request.Body)

	if err != nil {
		panic(err)
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(bytes.NewBuffer([]byte(res)).Bytes())
}
