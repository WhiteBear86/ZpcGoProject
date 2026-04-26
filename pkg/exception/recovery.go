package exception

import (
	"ZpcGoProject/pkg/resp"
	"github.com/gin-gonic/gin"
)

func GlobalException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp.Error(c, "服务器异常")
				c.Abort()
			}
		}()
		c.Next()
	}
}
