package repository

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"go-delic-products/model"
	"io"
)

type PostsRepo interface {
	/**
	Return the id of the document indexed
	*/
	Save(post *model.Post, c elasticsearch.Client) (string, error)

	/**
	In this example, I use the client and return the response from the elastic-go lib
	*/
	FindById(id string, c elasticsearch.Client) (*esapi.Response, error)

	/**
	Here I use the web client from Go, the query DSL coming from the client - perhaps is no the better idea
	but fits OK for an example, cURLs for free.
	*/
	FindByCriteria(criteria io.Reader) (string, error)
}
