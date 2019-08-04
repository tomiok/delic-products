package elastic

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/google/uuid"
	"go-delic-products/model"
	"strings"
)

type ElasticPost struct {
	repo elasticsearch.Client
}

func (p *ElasticPost) save(post *model.Post) (*model.Post, error) {
	jsonPost, _ := json.Marshal(post)
	request := esapi.IndexRequest{
		Index:      "posts",
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(string(jsonPost)),
		Refresh:    "true",
	}

}
