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
	Client elasticsearch.Client
}

func (p *PostElastic) Save(post *model.Post) (string, error) {

	jsonPost, _ := json.Marshal(post)
	request := esapi.IndexRequest{
		Index:      p.post.GetIndexName(),
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(string(jsonPost)),
		Refresh:    "true",
	}

	res, err := request.Do(context.Background(), &p.Client)

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

func (p *PostElastic) FindById(id string) (*esapi.Response, error) {
	req := esapi.GetRequest{Index: p.post.GetIndexName(), DocumentID: id}

	res, err := req.Do(context.Background(), &p.Client)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		log.Fatal("error parsing response")
	}

	return res, nil
}

func (p *PostElastic) FindByCriteria(criteria io.Reader) (string, error) {
	url := "http://localhost:9200/shared_post/_search"

	request, _ := ioutil.ReadAll(criteria)

	query := string(request)

	req, _ := http.NewRequest("POST", url, strings.NewReader(query))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}
