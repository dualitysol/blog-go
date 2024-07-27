package main

import (
	"github.com/dualitysol/go-crud/initializers"
	"github.com/dualitysol/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
