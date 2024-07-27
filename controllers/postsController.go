package controllers

import (
	"github.com/dualitysol/go-crud/services"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{
		postService: service,
	}
}

func (p *PostController) Create(c *gin.Context) {

}

func (p *PostController) Create(c *gin.Context) {

}

func (p *PostController) Create(c *gin.Context) {

}

func (p *PostController) Create(c *gin.Context) {

}

func (p *PostController) Create(c *gin.Context) {

}

func (p *PostController) Create(c *gin.Context) {

}
