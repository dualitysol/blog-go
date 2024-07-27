package repository

import (
	"errors"

	"github.com/dualitysol/go-crud/dto/posts/requests"
	"github.com/dualitysol/go-crud/helpers"
	"github.com/dualitysol/go-crud/models"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepositoryImpl(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{DB: db}
}

func (p *PostRepositoryImpl) CreateOne(post models.Post) (posts []models.Post, err error) {
	panic("unimplemented")
}

func (p *PostRepositoryImpl) FindAll() (posts []models.Post, err error) {
	var posts []models.Post

	r := p.DB.Find(&posts)

	if r != nil {
		return posts, nil
	} else {
		return posts, errors.New("Posts is not found!")
	}
}

func (p *PostRepositoryImpl) FindById(id uint) (post models.Post, err error) {
	post = models.Post{ID: id}

	r := p.DB.First(&post)

	if r != nil {
		return post, nil
	} else {
		return post, errors.New("Post is not found!")
	}
}

func (p *PostRepositoryImpl) Update(post models.Post) {
	var updatePost = requests.UpdatePost{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		AuthorId:    post.AuthorId,
		Media:       post.Media,
	}

	r := p.DB.Model(&post).Updates(updatePost)

	helpers.PanicError(r.Error)

}

func (p *PostRepositoryImpl) DeleteById(id uint) {
	r := p.DB.Where("id = ?", id).Delete(&models.Post)

	helpers.PanicError(r.Error)
}
