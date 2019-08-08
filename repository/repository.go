package repository

import (
	"github.com/elastic/go-elasticsearch"
	"go-delic-products/model"
)

type PostsRepo interface {
	Save(post *model.Post, c elasticsearch.Client) (string, error)

	FindById(id string, c elasticsearch.Client) (*model.Post, error)

	FindByCriteria(criteria string, c elasticsearch.Client) (*model.Post, error)
}
