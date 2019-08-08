package web

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
	"go-delic-products/elastic"
	"go-delic-products/model"
	"io/ioutil"
	"net/http"
)

func NewElasticWebHandler(c *elasticsearch.Client) *httpElastic {
	return &httpElastic{
		client: &elastic.PostElastic{
			Repo: c,
		},
	}
}

type httpElastic struct {
	client *elastic.PostElastic
}

func (esHandler *httpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)

	postApi := esHandler.client

	Repo := postApi.Repo

	idSaved, _ := postApi.Save(&post, *Repo)

	res, _ := json.Marshal(&idSaved)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (esHandler *httpElastic) GetByIdHandler(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]


	client := esHandler.client

	postApi := client.Repo
}
