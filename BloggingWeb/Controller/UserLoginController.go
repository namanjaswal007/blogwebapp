package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	models "BloggingWeb/Model"
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
	if !IsUserAuthenticated(map[string]string{"new_password": userCred.Password, "saved_password": user.Password}) {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1034 " + config.Message["UserPasswordWrongMsg"]})
		return
	}
	// Create user session
	if errDtls := UpdateUserSession(c, database, userCred.Email); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	if err = database.UpdateUserStatus(true, user.Uid); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1035" + config.Message["UserStatusUpdatingErrMsg"], Error: err})
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "user logged in successfully"})
}

func IsUserAuthenticated(userCredMap map[string]string) bool {
	var NewPassword, SavedPassword string
	var err error
	if userCredMap["new_password"] != "" {
		NewPassword, err = config.Decrypt(userCredMap["new_password"])
		if err != nil {
			return false
		}
	}
	if userCredMap["saved_password"] != "" {
		SavedPassword, err = config.Decrypt(userCredMap["saved_password"])
		if err != nil {
			return false
		}
	}
	return NewPassword == SavedPassword
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
	// Remove this code after testing, Because we will already get encrypted password from payload.
	err = config.Encrypt(&userCred.Password)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1039" + config.Message["EncryptionErrMsg"], Error: err})
		return
	}
	var user = view.User{FirstName: userCred.FirstName, LastName: userCred.LastName, FullName: userCred.FirstName + " " + userCred.LastName, Email: userCred.Email, Role: userCred.Role, Status: true}
	if err = database.AddUserDetails(&user); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1040" + config.Message[""], Error: err})
		return
	}
	userCred.Uid = user.ID
	database.SaveUserCredentials(&userCred)
	// This function is used to create user session like paseto token and save data into db.
	if errDtls := CreateUserSession(c, database, userCred); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "user registered successfully"})
}

func LogOutUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.DefaultQuery("uid", ""))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1040" + config.Message[""], Error: err})
		return
	}
	database, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1037" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer config.DisconnectDbConnection(database.MainDB)
	if err := UpdateUserStatus(database, uid); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1041" + config.Message[""], Error: err})
		return
	}
	config.GetErrorResponse(c, view.SuccessResp{Response: "user logout successfully"})
}

func UpdateUserStatus(db models.Database, uid int) (err error) {
	if err = db.UpdateUserSessionStatus(uid); err != nil {
		return
	}
	if err = db.UpdateUserStatus(false, uid); err != nil {
		return
	}
	return
}
