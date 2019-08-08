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
	Repo *elasticsearch.Client
	post model.Post
}

func (p *PostElastic) Save(post *model.Post, c elasticsearch.Client) (string, error) {

	jsonPost, _ := json.Marshal(post)
	request := esapi.IndexRequest{
		Index:      p.post.Content,
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(string(jsonPost)),
		Refresh:    "true",
	}

	res, err := request.Do(context.Background(), &c)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return "", errors.New("errors during the response")
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		}
		return r["_id"].(string), nil
	}

}

func (p *PostElastic) FindById(id string, c elasticsearch.Client) (*model.Post, error) {
	req := esapi.GetRequest{Index: "", DocumentID: id}
	return nil, nil
}

func (p *PostElastic) FindByCriteria(id string, c elasticsearch.Client) (*model.Post, error) {
	return nil, nil
}
