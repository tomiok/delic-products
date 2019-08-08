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

func (es *httpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &post)

	client := elastic.PostElastic{}

	Repo := es.client.Repo

	idSaved, _ := client.Save(&post, *Repo)

	res, _ := json.Marshal(&idSaved)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func GetByIdHandler(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]


}
