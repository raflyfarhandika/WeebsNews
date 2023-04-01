package middleware

import (
	"net/http"
	"weebsnews/helper"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, please login first")
			c.Abort()
			return
		}
		c.Next()
	}
}