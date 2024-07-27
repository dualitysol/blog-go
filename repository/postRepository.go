package repository

import (
	"github.com/dualitysol/go-crud/dto/posts/requests"
	"github.com/dualitysol/go-crud/models"
)

type PostRepository interface {
	CreateOne(post models.Post) (posts []models.Post, err error)
	FindAll() (post []models.Post, err error)
	FindById(id uint) (post models.Post, err error)
	UpdateById(p requests.UpdatePost) (post models.Post, err error)
	DeleteById(id uint) (success bool, err error)
}
