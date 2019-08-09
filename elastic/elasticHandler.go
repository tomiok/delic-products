package elastic

import (
	"bytes"
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
	post model.Post
}

func (p *PostElastic) Save(post *model.Post, c elasticsearch.Client) (string, error) {

	jsonPost, _ := json.Marshal(post)
	request := esapi.IndexRequest{
		Index:      p.post.GetIndexName(),
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

func (p *PostElastic) FindById(id string, c elasticsearch.Client) (*esapi.Response, error) {
	req := esapi.GetRequest{Index: p.post.GetIndexName(), DocumentID: id}

	res, err := req.Do(context.Background(), &c)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		log.Fatal("error parsing response")
	}

	return res, nil
}

func (p *PostElastic) FindByCriteria(criteria string, c elasticsearch.Client) (*esapi.Response, error) {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(criteria)
	res , err := c.Search(
		c.Search.WithContext(context.Background()),
		c.Search.WithBody(&buf),
		c.Search.WithPretty(),
		c.Search.WithTrackTotalHits(true),
		c.Search.WithIndex("shared_post"),
	)

	if err != nil {
		log.Fatal("errors during the request", err)
		return nil, err
	}

	if res.IsError() {
		log.Fatal("errors parsing the response")
	}

	return res, nil
}
