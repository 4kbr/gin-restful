package main

import (
	"fmt"
	"gin-restful/controller"
	"gin-restful/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controller.PostsCreate)
	r.PUT("/posts/:id", controller.PostsUpdateById)

	r.GET("/posts", controller.PostsGetAll)
	r.GET("/posts/:id", controller.PostsGetById)

	r.DELETE("/posts/:id", controller.PostsRemoveById)

	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("running server")
}
