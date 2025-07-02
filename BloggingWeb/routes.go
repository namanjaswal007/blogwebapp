package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/blog-app/controllers"

)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.POST("/", controllers.CreatePost)
		// post.GET("/", controllers.GetPosts)
		// post.GET("/:id", controllers.GetPostByID)
		// post.PUT("/:id", controllers.UpdatePost)
		// post.DELETE("/:id", controllers.DeletePost)
	}
	r.Run(":8080")
}
