package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	PostRoutes(r)

}
