package repository

import (
	"gin-restful/helper"
	"gin-restful/initializers"
	"gin-restful/model"
)

func PostSelectAll(posts *[]model.Post) error {
	result := initializers.DB.Raw("SELECT * FROM posts").Scan(&posts)
	return helper.ResultValidation(result.Error)
}

func PostSelectById(id int, post *model.Post) error {
	result := initializers.DB.Raw("SELECT * FROM posts WHERE id = ?;", id).Scan(&post)
	return helper.ResultValidation(result.Error)
}

func PostUpdate(body struct {
	Body  string
	Title string
}, post *model.Post) error {
	result := initializers.DB.Model(&post).Updates(model.Post{
		Body:  body.Body,
		Title: body.Title,
	})

	return helper.ResultValidation(result.Error)
}

func PostDeleteByid(id int) error {

	result := initializers.DB.Delete(&model.Post{}, "5")

	return helper.ResultValidation(result.Error)
}
