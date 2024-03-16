package controller

import (
	"gin-restful/initializers"
	"gin-restful/model"
	"gin-restful/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//TODO: get request
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	//TODO: create a post
	post := model.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//TODO: return
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsGetAll(c *gin.Context) {
	var posts []model.Post

	repository.PostSelectAll(&posts)

	c.JSON(200, gin.H{
		"data": &posts,
	})
}

func PostsGetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var post model.Post
	repository.PostSelectById(id, &post)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsUpdateById(c *gin.Context) {
	// Cari post berdasarkan ID
	id, _ := strconv.Atoi(c.Param("id"))

	//TODO: get request
	var body struct {
		Body  string
		Title string
	}
	var post model.Post

	repository.PostSelectById(id, &post)
	// result := db.Raw("SELECT * FROM posts WHERE id = ?", id).Scan(&post)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Bind data dari request body ke struct Post
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.PostUpdate(body, &post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
func PostsRemoveById(c *gin.Context) {
	// Cari post berdasarkan ID
	id, _ := strconv.Atoi(c.Param("id"))

	var post model.Post

	err := repository.PostSelectById(id, &post)
	if err != nil || post.DeletedAt.Valid {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	err = repository.PostDeleteByid(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete post " + strconv.Itoa(id),
	})
}
