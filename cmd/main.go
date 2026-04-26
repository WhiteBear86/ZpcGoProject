package main

import (
	"ZpcGoProject/config"
	"ZpcGoProject/router"
	"fmt"
)

func main() {
	config.InitNacosConfig()
	config.InitDB()
	config.RegisterService()

	r := router.Init()
	fmt.Println("🚀 服务启动成功 : " + config.ServerPort)
	r.Run(":" + config.ServerPort)
}
