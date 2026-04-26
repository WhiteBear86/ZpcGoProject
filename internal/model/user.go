package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=18"`
}
