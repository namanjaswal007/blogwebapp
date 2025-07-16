package controller

import (
	"BloggingWeb/BlogDB"
	config "BloggingWeb/Config"
	models "BloggingWeb/Model"
	view "BloggingWeb/View"
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

)

func UserLogin(c *gin.Context) {
	var userCred view.UserCredentials
	if err := c.ShouldBindJSON(&userCred); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1035" + config.Message["ErrorWhileReadingPayloadData"], Error: err})
		return
	}
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1032" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer db.MainDB.Close()
	db.Query = BlogDB.New(db.MainDB)
	user, err := db.Query.GetUserCredByEmail(context.Background(), userCred.Email)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1032" + config.Message["UserSessionErrMsg"], Error: err})
		return
	}
	// In this function, Check user the Authenticated or not by decrypting user saved password and provided password, After comparing th both password are same or not.
	if !IsUserAuthenticated(map[string]string{"new_password": userCred.Password, "saved_password": user.Password}) {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1034 " + config.Message["UserPasswordWrongMsg"]})
		return
	}
	// HandleUserLoginState update a session for the user and updates their login status.
	if errDtls := HandleUserLoginState(c, db, userCred.Email, int(user.Uid)); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: "user logged in successfully"})
}

// HandleUserLoginState update a session for the user and updates their login status.
func HandleUserLoginState(c *gin.Context, db models.Database, email string, uid int) (errDtls view.ErrResp) {
	if errDtls = UpdateUserSession(c, db, email); errDtls.Error != nil {
		config.GetErrorResponse(c, errDtls)
		return
	}
	if errDtls.Error = db.Query.UpdateUserStatus(context.Background(), BlogDB.UpdateUserStatusParams{UserID: int64(uid), Status: sql.NullBool{Bool: true, Valid: true}}); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1035" + config.Message["UserStatusUpdatingErrMsg"]
	}
	return
}

// In this function, compares the provided and stored passwords after decryption.
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
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1037 " + config.Message["ErrorWhileConnectingDB"], Error: err.Error()})
		return
	}
	defer db.MainDB.Close()
	db.Query = BlogDB.New(db.MainDB)
	exist, err := db.Query.CheckUserRegistration(context.Background(), userCred.Email)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1038 " + config.Message["CheckUserRegErrMsg"], Error: err})
		return
	}
	successMsg := "This email is already registered. Please try another one."
	if !exist {
		if errDtls := handleUserRegistration(c, db, userCred); errDtls.Error != nil {
			config.GetErrorResponse(c, errDtls)
			return
		}
		successMsg = "User register successfully"
	}
	config.GetSuccessResponse(c, view.SuccessResp{SuccessMsg: successMsg})
}

func handleUserRegistration(c *gin.Context, db models.Database, userCred view.UserCredentials) (errDtls view.ErrResp) {
	// Remove this code after testing, Because we will already get encrypted password from payload.
	if errDtls.Error = config.Encrypt(&userCred.Password); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1039" + config.Message["EncryptionErrMsg"]
		return
	}
	// CompleteUserRegistration handles the full user registration process.
	// It saves user profile details and login credentials to the database,
	// generates a PASETO token, and creates a session for the user.
	errDtls = SaveUserRegistrationWithSession(c, db, userCred)
	return
}

// CompleteUserRegistration handles the full user registration process.
// It saves user profile details and login credentials to the database,
// generates a PASETO token, and creates a session for the user.
func SaveUserRegistrationWithSession(c *gin.Context, db models.Database, userCred view.UserCredentials) (errDtls view.ErrResp) {
	var user = view.User{FirstName: userCred.FirstName, LastName: userCred.LastName, FullName: userCred.FirstName + " " + userCred.LastName, Email: userCred.Email, Role: userCred.Role, Status: true}
	if errDtls.Error = db.AddUserDetails(&user); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1040" + config.Message["SavingUserDetailsErrMsg"]
		return
	}
	userCred.Uid = user.UserId
	fmt.Println("User details : ", user)
	if errDtls.Error = db.Query.SaveUserCredentials(context.Background(), BlogDB.SaveUserCredentialsParams{FirstName: userCred.FirstName, LastName: userCred.LastName, Email: userCred.Email, Password: userCred.Password, Uid: int64(user.UserId), Role: user.Role}); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1041" + config.Message["SavingUserCredentialsErrMsg"]
		return
	}
	// This function is used to create user session like paseto token and save data into db.
	errDtls = CreateUserSession(c, db, userCred)
	return
}

func LogOutUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.DefaultQuery("uid", ""))
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1040" + config.Message[""], Error: err})
		return
	}
	db, err := config.ConnectDB(config.MainDB)
	if err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1037" + config.Message["ErrorWhileConnectingDB"], Error: err})
		return
	}
	defer db.MainDB.Close()
	db.Query = BlogDB.New(db.MainDB)
	if err := UpdateUserStatus(db, uid); err != nil {
		config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error #1041" + config.Message[""], Error: err})
		return
	}
	config.GetErrorResponse(c, view.SuccessResp{Response: "user logout successfully"})
}

func UpdateUserStatus(db models.Database, uid int) (err error) {
	if err = db.Query.UpdateUserSessionStatus(context.Background(), BlogDB.UpdateUserSessionStatusParams{Uid: int64(uid), Status: false}); err != nil {
		return
	}
	err = db.Query.UpdateUserStatus(context.Background(), BlogDB.UpdateUserStatusParams{UserID: int64(uid), Status: sql.NullBool{Bool: false, Valid: true}})
	return
}
