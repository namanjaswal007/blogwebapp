package controller

import (
	"BloggingWeb/BlogDB"
	config "BloggingWeb/Config"
	view "BloggingWeb/View"
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	defer database.MainDB.Close()
	query := BlogDB.New(database.MainDB)
	dob, errDtls := config.ConvStrToTimeStamp(c, user.DateOfBirth)
	if errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	updatedUserDtls, err := query.UpdateUserDetails(context.Background(), BlogDB.UpdateUserDetailsParams{UserID: int64(uid), FirstName: user.FirstName, LastName: user.LastName, FullName: user.FullName, DateOfBirth: sql.NullTime{Time: dob, Valid: true}, UserAddress: sql.NullString{String: user.UserAddress, Valid: true}, UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1022" + config.Message["ErrorWhileCheckUser"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: updatedUserDtls})
}

func GetAllUsers(c *gin.Context) {
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1023" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer database.MainDB.Close()
	var users []BlogDB.User
	query := BlogDB.New(database.MainDB)
	if users, err = query.GetAllUsers(context.Background()); err != nil {
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
	defer db.MainDB.Close()
	query := BlogDB.New(db.MainDB)
	var user BlogDB.User
	if user, err = query.GetUserByUid(context.Background(), int64(uid)); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1027" + config.Message["ErrorWhileGettingUserById"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "Done", Response: user})
}
