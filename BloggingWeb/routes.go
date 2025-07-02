package main

import (
	"github.com/gin-gonic/gin"

	controler "BloggingWeb/Controler"

)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.POST("create_post", controler.CreatePost)
		post.GET("/get_post", controler.GetPosts)
		// post.DELETE("/:id", controllers.DeletePost)
	}
	r.Run(":8080")
}
