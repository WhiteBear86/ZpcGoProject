package config

import (
	"ZpcGoProject/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
*
初始化gorm数据库连接源的函数
*/
func InitDB() {
	db, err := gorm.Open(mysql.Open(DBConfig.DSN()), &gorm.Config{})
	if err != nil {
		panic("MySQL connect failed: " + err.Error())
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic("MySQL migration failed: " + err.Error())
	}

	DB = db
}
