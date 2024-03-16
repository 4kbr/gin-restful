package main

import (
	"gin-restful/initializers"
	"gin-restful/model"
)

func init() {
	// initializers.
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
