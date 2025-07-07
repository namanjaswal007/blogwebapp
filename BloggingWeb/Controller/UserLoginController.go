package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	view "BloggingWeb/View"
)

func UserLogin(c *gin.Context) {
	var userCred view.UserCredentials
	if err := c.ShouldBindJSON(&userCred); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1035" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1032" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	user, err := database.GetUserCredByEmail(userCred.Email)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1032" + config.Message["UserSessionErrMsg"], Error: err})
		return
	}
	decryptedPassword, err := config.Decrypt(user.Password)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #10333" + config.Message["DecryptionErrMsg"], Error: err})
		return
	}
	if userCred.Password != decryptedPassword {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1034" + config.Message["UserPasswordWrongMsg"]})
		return
	}
	// Create user session
	if errDtls := UpdateUserSession(c, database, userCred.Email); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "user logged in successfully"})
}

func UserRegister(c *gin.Context) {
	var userCred view.UserCredentials
	if err := c.ShouldBindJSON(&userCred); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1036" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}

	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1037" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	_, err = database.GetUserCredByEmail(userCred.Email)
	if err == nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1038 user is already existed", Error: err})
		return
	}
	// err = config.Encrypt(&userCred.Password)
	// if err != nil {
	// 	config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1039" + config.Message["EncryptionErrMsg"], Error: err})
	// 	return
	// }
	database.SaveUserCredentials(&userCred)
	fmt.Println(userCred)
	if errDtls := CreateUserSession(c, database, userCred); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}

	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "user registered successfully"})
}
