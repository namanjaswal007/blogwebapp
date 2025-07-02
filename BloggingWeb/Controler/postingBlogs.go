package controler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	model "BloggingWeb/Model"
	view "BloggingWeb/View"
)

func CreatePost(c *gin.Context) {
	var post view.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access DB"})
		return
	}
	if err := model.CreatePost(database, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	var posts []view.Post
	database, err := config.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access DB"})
		return
	}
	if err := model.GetAllPosts(database, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
