package main

import (
	"github.com/gin-gonic/gin"

	controller "BloggingWeb/Controller"
	middleware "BloggingWeb/Middleware"

)

func PostRoutes(r *gin.Engine) {
	public := r.Group("/posts")
	{
		public.POST("create_post", controller.CreateBlog)
		public.POST("update_blog_content", controller.UpdateBlog)
		public.GET("/get_blogs", controller.GetBlogs)
		public.GET("/get_blog/:id", controller.GetBlogByID)
		public.POST("/add_user", controller.AddUser)
		public.GET("get_user_blogs", controller.GetUserBlogs)
		public.GET("login", controller.Login)
	}
	private := r.Group("/api")
	{
		private.Use(middleware.AuthPasetoMiddleware())
		private.GET("get_users", controller.GetAllUsers)
		private.GET("get_user/:id", controller.GetUserByUid)
		private.DELETE("/delete_blog/:id", controller.DeleteBlog)
	}

	r.Run(":8080")
}
