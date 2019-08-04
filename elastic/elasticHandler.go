package elastic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/google/uuid"
	"go-delic-products/model"
	"log"
	"strings"
)

type PostElastic struct {
	Repo elasticsearch.Client
}

func (p *PostElastic) Save(post *model.Post, client *elasticsearch.Client) (*model.Post, error) {
	jsonPost, _ := json.Marshal(post)
	request := esapi.IndexRequest{
		Index:      "posts",
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(string(jsonPost)),
		Refresh:    "true",
	}

	res, err := request.Do(context.Background(), client)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return &model.Post{}, errors.New("errors during the response")
	} else {
		var r = model.Post{}

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		}
		return &r, nil
	}

}

func (p *PostElastic) FindById(id string) (*model.Post, error) {
	return nil, nil
}

func (p *PostElastic) FindByCriteria(id string) (*model.Post, error) {
	return nil, nil
}
