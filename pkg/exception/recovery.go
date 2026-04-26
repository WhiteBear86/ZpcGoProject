package exception

import (
	"ZpcGoProject/pkg/resp"
	"github.com/gin-gonic/gin"
)

func GlobalException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp.Error(c, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
