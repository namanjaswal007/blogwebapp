package controler

import (
	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	model "BloggingWeb/Model"
	view "BloggingWeb/View"
)

func CreatePost(c *gin.Context) {
	var post view.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		config.GetErrorResponse(c, err)
		return
	}
	database, err := config.ConnectDB()
	if err != nil {
		config.GetErrorResponse(c, err)
		return
	}
	if err := model.CreatePost(database, &post); err != nil {
		config.GetErrorResponse(c, err)
		return
	}
	config.GetSuccessResponse(c, post)
}

func GetPosts(c *gin.Context) {
	var posts []view.Post
	database, err := config.ConnectDB()
	if err != nil {
		config.GetErrorResponse(c, err)
		return
	}
	if err := model.GetAllPosts(database, &posts); err != nil {
		config.GetErrorResponse(c, err)
		return
	}
	config.GetSuccessResponse(c, posts)
}
