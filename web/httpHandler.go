package web

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"go-delic-products/elastic"
	"go-delic-products/model"
	"net/http"
)

func newElasticWebHandler(c elasticsearch.Client) *HttpElastic {
	return &HttpElastic{
		client: elastic.PostElastic{
			Repo: c,
		},
	}
}

type HttpElastic struct {
	client elastic.PostElastic
}

func (es *HttpElastic) SaveHandler(w http.ResponseWriter, r *http.Request) {
	post := model.Post{}

	_ = json.NewDecoder(r.Body).Decode(&post)

	res, _ := json.Marshal(post)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}
