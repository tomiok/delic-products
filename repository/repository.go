package repository

import (
	"github.com/elastic/go-elasticsearch"
	"go-delic-products/model"
)

type PostsRepo interface {
	Save(post *model.Post, client *elasticsearch.Client) (*model.Post, error)

	FindById(id string) (*model.Post, error)

	FindByCriteria(criteria string) (*model.Post, error)
}
