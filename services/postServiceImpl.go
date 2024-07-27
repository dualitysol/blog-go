package services

import (
	"github.com/dualitysol/go-crud/dto/posts/requests"
	"github.com/dualitysol/go-crud/dto/posts/response"
	"github.com/dualitysol/go-crud/helpers"
	"github.com/dualitysol/go-crud/models"
	"github.com/dualitysol/go-crud/repository"
	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	validate       *validator.Validate
}

// Create implements PostService.
func (p *PostServiceImpl) Create(post requests.CreatePost) {
	err := p.validate.Struct(post)

	helpers.PanicError(err)

	postModel := models.Post{
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		AuthorId:    post.AuthorId,
		Media:       post.Media,
	}

	p.PostRepository.CreateOne(postModel)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(id uint) {
	p.PostRepository.DeleteById(id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll() []response.PostResult {
	result, err := p.PostRepository.FindAll()

	helpers.PanicError(err)

	var posts []response.PostResult

	for _, value := range result {
		post := response.PostResult{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			Content:     value.Content,
			AuthorId:    value.AuthorId,
			Media:       value.Media,
		}

		posts = append(posts, post)
	}

	return posts
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(id uint) response.PostResult {
	postData, err := p.PostRepository.FindById(id)

	helpers.PanicError(err)

	postResponse := response.PostResult{
		ID:          postData.ID,
		Title:       postData.Title,
		Description: postData.Description,
		Content:     postData.Content,
		AuthorId:    postData.AuthorId,
		Media:       postData.Media,
	}

	return postResponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(post requests.UpdatePost) {
	postData, err := p.PostRepository.FindById(post.ID)

	helpers.PanicError(err)

	postData.Title = post.Title
	postData.Description = post.Description
	postData.Content = post.Content
	postData.AuthorId = post.AuthorId
	postData.Media = post.Media

	p.PostRepository.UpdateById(post)
}

func NewPostServiceImpl(postRepository repository.PostRepository, validator validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		validate:       &validator,
	}
}
