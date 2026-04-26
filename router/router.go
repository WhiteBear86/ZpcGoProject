package router

import (
	"ZpcGoProject/internal/handler"
	"ZpcGoProject/pkg/exception"
	"ZpcGoProject/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(exception.GlobalException())

	g := r.Group("/api/user")
	g.Use(middleware.ApiTokenAuth())
	{
		g.POST("", handler.Add)
		g.GET("/:id", handler.Get)
		g.GET("/list", handler.List)
		g.PUT("", handler.Update)
		g.DELETE("/:id", handler.Delete)
	}
	return r
}
