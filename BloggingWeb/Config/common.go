package config

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	view "BloggingWeb/View"

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

func ConvStrToTimeStamp(c *gin.Context, date string) (dob time.Time, errDtls view.ErrResp) {
	dob, errDtls.Error = time.Parse("2006-01-02", date)
	if errDtls.Error != nil {
		errDtls.ErrMsg = "Invalid date format. Use YYYY-MM-DD"
		return
	}
	return
}
