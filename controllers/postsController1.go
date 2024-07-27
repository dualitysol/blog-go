package controllers

import (
	"strconv"

	"github.com/dualitysol/go-crud/initializers"
	"github.com/dualitysol/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get body props from request
	var body struct {
		Title       string
		Description string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Description: body.Description}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)

		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostGetAll(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostGetById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdateById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 19)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Invalid post id!",
		})
		return
	}

	var post models.Post

	c.Bind(&post)

	post.ID = uint(id)

	initializers.DB.Model(&post).Update("name", "hello")

	// log.Fatalln(tx)
	// if tx.Error != nil {
	// 	fmt.Println("ERROR", tx.Error)
	// 	c.JSON(500, gin.H{
	// 		"error": tx.Error,
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"post": post,
	})
}
