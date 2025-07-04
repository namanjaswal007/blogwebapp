package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	middleware "BloggingWeb/Middleware"
	view "BloggingWeb/View"

)

func Login(c *gin.Context) {
	// if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	// 	return
	// }
	reqData := view.RequestData{
		Username: "Mayank",
		Role:     "Merchant",
		Email:    "jaswalmayank0@gmail.com",
		Uid:      101,
		Password: "jaswal-1730",	
	}
	fmt.Println(reqData)
	token, err := middleware.CreateToken(reqData, time.Hour)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		fmt.Println("Error while creating token", err)
	}
	fmt.Println("Token :", token)
}
