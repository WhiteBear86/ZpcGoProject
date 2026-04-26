package config

import (
	"ZpcGoProject/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库对象
var DB *gorm.DB

// 初始化数据库
func InitDB() {
	// 这里必须用 DBConfig.xxx ！！！
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.Username,
		DBConfig.Password,
		DBConfig.Host,
		DBConfig.Port,
		DBConfig.Database,
		DBConfig.Charset,
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("MySQL 连接失败：" + err.Error())
	}

	// 自动建表
	db.AutoMigrate(&model.User{})

	// 赋值给全局 DB
	DB = db
}
