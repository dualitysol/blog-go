package main

import (
	"github.com/dualitysol/go-crud/controllers"
	"github.com/dualitysol/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.PostGetAll)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.PostGetById)
	r.PUT("/posts/:id", controllers.PostUpdateById)

	r.Run()
}
