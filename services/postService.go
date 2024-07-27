package services

import (
	"github.com/dualitysol/go-crud/dto/posts/requests"
	"github.com/dualitysol/go-crud/dto/posts/response"
)

type PostService interface {
	Create(post requests.CreatePost)
	Update(post requests.UpdatePost)
	Delete(id uint)
	FindById(id uint) response.PostResult
	FindAll() []response.PostResult
}
