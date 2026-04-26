package handler

import (
	"strconv"

	"ZpcGoProject/internal/model"
	"ZpcGoProject/internal/service"
	"ZpcGoProject/pkg/resp"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)
	if err := service.Create(&u); err != nil {
		resp.Error(c, "create user failed")
		return
	}
	resp.Success(c, u)
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := service.GetByID(uint(id))
	if err != nil {
		resp.Error(c, "user not found")
		return
	}
	resp.Success(c, user)
}

func List(c *gin.Context) {
	list, _ := service.List()
	resp.Success(c, list)
}

func Update(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)
	_ = service.Update(&u)
	resp.Success(c, "update success")
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = service.Delete(uint(id))
	resp.Success(c, "delete success")
}
