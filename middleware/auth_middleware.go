package middleware

import (
	"net/http"
	"strings"
	"test_kredit_plus/helper"

	"github.com/gin-gonic/gin"
)

func JWTChatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		tokenStringHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		claims, err := helper.ValidateJWT(tokenStringHeader)
		_, err = helper.ValidateChatJWT(tokenStringHeader)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   true,
			})
			c.Abort()
			return
		}

		//  test
		c.Set("username", claims["username"])
		c.Set("user", claims["user"])
		c.Next()
	}
}
