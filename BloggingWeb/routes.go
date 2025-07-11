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
		public.GET("/get_blogs", controller.GetBlogs)
		public.GET("/get_blog/:id", controller.GetBlogByID)
		public.GET("get_user_blogs", controller.GetUserBlogs)
		public.GET("user_login", controller.UserLogin)
		public.POST("user_register", controller.UserRegister)
	}
	private := r.Group("/api")
	private.Use(middleware.AuthPasetoMiddleware())
	{
		private.POST("/update_user", controller.UpdateUser)
		private.POST("update_blog_content", controller.UpdateBlog)
		private.GET("get_users", controller.GetAllUsers)
		private.GET("get_user/:id", controller.GetUserByUid)
		private.DELETE("/delete_blog/:id", controller.DeleteBlog)
		private.GET("/logout", controller.LogOutUser)
	}
	r.Run(":8080")
}
