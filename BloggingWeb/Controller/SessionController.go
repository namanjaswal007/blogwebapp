package controller

import (
	"BloggingWeb/BlogDB"
	config "BloggingWeb/Config"
	middleware "BloggingWeb/Middleware"
	models "BloggingWeb/Model"
	view "BloggingWeb/View"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUserSession(c *gin.Context, db models.Database, userCred view.UserCredentials) (errDtls view.ErrResp) {
	var token string
	token, errDtls.Error = middleware.CreateUserSessionToken(view.UserSessionToken{Username: userCred.FirstName + " " + userCred.LastName, Email: userCred.Email, Role: userCred.Role}, time.Hour)
	if errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1042: " + config.Message["PasetoTokenErrMsg"]
		return
	}
	session := BlogDB.SaveUserSessionParams{
		Email:     userCred.Email,
		Role:      userCred.Role,
		UserAgent: c.GetHeader("user-Agent"),
		Token:     token,
		Uid:       int64(userCred.Uid),
	}
	if errDtls.Error = db.Query.SaveUserSession(context.Background(), session); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1043: " + config.Message["SaveSessionErrMsg"]
	}
	return
}

func UpdateUserSession(c *gin.Context, db models.Database, email string) (errDtls view.ErrResp) {
	// Get user from DB
	var user BlogDB.User
	if user, errDtls.Error = db.Query.GetUserByEmail(context.Background(), email); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1044: " + config.Message["GetUserErrMsg"]
		return
	}
	// Create new token
	newToken, err := middleware.CreateUserSessionToken(view.UserSessionToken{Username: user.FullName, Email: user.Email, Role: user.Role.String}, time.Hour)
	if err != nil {
		errDtls.ErrMsg = "Error #1045: " + config.Message["PasetoTokenErrMsg"]
		errDtls.Error = err
		return
	}
	// userAgent := c.GetHeader("User-Agent")
	if _, errDtls.Error = db.Query.UpdateSessionTokenAndAgent(context.Background(), BlogDB.UpdateSessionTokenAndAgentParams{Token: newToken, UserAgent: c.GetHeader("User-Agent"), CreatedAt: time.Now(), Status: true, Email: email}); errDtls.Error != nil {
		errDtls.ErrMsg = "Error #1046: " + config.Message["UpdateSessionErrMsg"]
	}
	return
}
