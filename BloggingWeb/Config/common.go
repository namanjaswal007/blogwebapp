package config

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

func GetErrorResponse(c *gin.Context, errMsg interface{}) {
	c.JSON(http.StatusInternalServerError, errMsg)
}

func GetSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func DisconnectDbConnection(database *gorm.DB) {
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
sqlDB.Close()
}
