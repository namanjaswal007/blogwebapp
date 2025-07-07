package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"

	config "BloggingWeb/Config"
	view "BloggingWeb/View"
)

func CreateUserSessionToken(data view.UserSessionToken, duration time.Duration) (string, error) {
	data.Exp = time.Now().Add(duration)
	return paseto.NewV2().Encrypt([]byte(config.SymmetricKey), data, nil)
}

func AuthPasetoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error: #1030 " + config.Message["PesatoAuthTokenErrMsg"]})
			return
		}
		payload, err := VerifyToken(token)
		if err != nil {
			config.GetErrorResponse(c, view.ErrResp{ErrMsg: "Error: #1031 " + config.Message["InvalidToken"], Error: err})
			return
		}
		c.Set("username", payload.Username)
		c.Set("email", payload.Email)
		c.Set("role", payload.Role)
		c.Next()
	}
}

func VerifyToken(token string) (payload view.UserSessionToken, err error) {
	err = paseto.NewV2().Decrypt(token, []byte(config.SymmetricKey), &payload, nil)
	if err != nil {
		return
	}
	if time.Now().After(payload.Exp) {
		err = fmt.Errorf("token expired")
		return
	}
	return
}
