package repository

import "go-delic-products/model"

type PostsRepo interface {

	Save(post *model.Post) (*model.Post, error)

	FindById(id string) (*model.Post, error)

	FindByCriteria(criteria string) (*model.Post, error)
}
