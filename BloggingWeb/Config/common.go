package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func GetErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err)
}

func GetSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}