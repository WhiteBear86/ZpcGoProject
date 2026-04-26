package middleware

import (
	"ZpcGoProject/config"
	"ZpcGoProject/pkg/resp"
	"github.com/gin-gonic/gin"
)

func ApiTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Api-Token")
		if token == "" || token != config.ApiToken {
			resp.Unauthorized(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
