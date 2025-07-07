package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	config "BloggingWeb/Config"
	middleware "BloggingWeb/Middleware"
	models "BloggingWeb/Model"
	view "BloggingWeb/View"

)

func CreateUserSession(c *gin.Context, db models.Database, userCred view.UserCredentials) (errDtls view.ErrResp) {
	var token string
	token, errDtls.Error = middleware.CreateUserSessionToken(view.UserSessionToken{Username: userCred.FirstName+" "+userCred.LastName, Email: userCred.Email, Role: userCred.Role}, time.Hour)
	if errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1042: " + config.Message["PasetoTokenErrMsg"]
		return
	}
	session := view.UserSession{
		Email:     userCred.Email,
		Role:      userCred.Role,
		UserAgent: c.GetHeader("user-Agent"),
		Token:     token,
		CreatedAt: time.Now(),
		ID: userCred.ID,
	}
	if errDtls.Error = db.SaveSession(&session); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1043: " + config.Message["SaveSessionErrMsg"]
		return
	}
	return
}

func UpdateUserSession(c *gin.Context, db models.Database, email string) (errDtls view.ErrResp) {
	// Get user from DB
	var user view.User
	if errDtls.Error = db.GetUserByEmail(&user, email); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1044: " + config.Message["GetUserErrMsg"]
		return
	}
	// Create new token
	newToken, err := middleware.CreateUserSessionToken(view.UserSessionToken{
		Username: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}, time.Hour)
	if err != nil {
		errDtls.ErrMsg = "Error #1045: " + config.Message["PasetoTokenErrMsg"]
		errDtls.Error = err
		return
	}
	userAgent := c.GetHeader("User-Agent")
	if errDtls.Error = db.UpdateSessionTokenAndAgent(view.UserSession{UserAgent: userAgent, Token: newToken, Email: email}); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1046: " + config.Message["UpdateSessionErrMsg"]
		return
	}

	return
}
