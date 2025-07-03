package main

import (
	"github.com/gin-gonic/gin"

	controller "BloggingWeb/Controller"

)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.POST("create_post", controller.CreateBlog)
		post.POST("update_blog_content", controller.UpdateBlog)
		post.GET("/get_blogs", controller.GetBlogs)
		post.DELETE("/delete_blog/:id", controller.DeleteBlog)
		post.GET("/get_blog/:id", controller.GetBlogByID)
		post.POST("/add_user", controller.AddUser)
		post.GET("get_users", controller.GetAllUsers)
		post.GET("get_user/:id", controller.GetUserByUid)
		post.GET("get_user_blogs", controller.GetUserBlogs)
	}
	r.Run(":8080")
}
