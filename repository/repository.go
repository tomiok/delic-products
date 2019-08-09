package repository

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"go-delic-products/model"
)

type PostsRepo interface {
	/**
	Return the id of the document indexed
	*/
	Save(post *model.Post, c elasticsearch.Client) (string, error)

	FindById(id string, c elasticsearch.Client) (*esapi.Response, error)

	FindByCriteria(criteria string, c elasticsearch.Client) (*esapi.Response, error)
}
