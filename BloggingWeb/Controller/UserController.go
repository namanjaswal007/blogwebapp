package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	view "BloggingWeb/View"
)

func UpdateUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.DefaultQuery("uid", ""))
	var user view.User
	if err := c.ShouldBindJSON(&user); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1020" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1021" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	err = database.UpdateUserDetails(map[string]interface{}{"date_of_birth": user.DateOfBirth, "user_address": user.UserAddress}, uid)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1022" + config.Message["ErrorWhileCheckUser"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done"})
}

func GetAllUsers(c *gin.Context) {
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1023" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	var users []view.User
	if err = database.GetAllUsers(&users); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1024" + config.Message["ErrorWhileGettingAllUsers"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: users})
}

func GetUserByUid(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1025" + config.Message["ErrorWhileConvertingStrToInt"], Error: err})
		return
	}
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1026" + config.Message["ErrorWhileGettingAllUsers"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(db.MainDB)
	var user view.User
	if err = db.GetUserByUid(&user, uid); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1027" + config.Message["ErrorWhileGettingUserById"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: user})
}
