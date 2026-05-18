package middleware

import (
	appErr "server/internal/shared/errors"
	"server/internal/shared/response"
	"server/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("access_token")
		if err != nil {
			response.HandleError(c, appErr.Unauthorized("missing access token", nil))
			c.Abort()
			return
		}

		claims, err := utils.VerifyAccessToken(cookie.Value, secret)
		if err != nil {
			response.HandleError(c, appErr.Unauthorized("invalid token", nil))
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)

		c.Next()
	}
}
