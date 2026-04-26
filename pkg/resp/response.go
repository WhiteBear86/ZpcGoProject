package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{200, "success", data})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Result{500, msg, nil})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Result{401, "token无效", nil})
}
