package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/blog-app/routes"
)

func main() {
	r := gin.Default()

	routes.PostRoutes(r)

}
