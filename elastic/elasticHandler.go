package elastic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/google/uuid"
	"go-delic-products/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func (p *PostElastic) FindByCriteria(criteria io.Reader) (string, error) {
	res, _ := http.Post("http://localhost:9200/shared_post/_search", "application/json", criteria)

	httpRes, _ := ioutil.ReadAll(res.Body)

	return string(httpRes), nil
}
