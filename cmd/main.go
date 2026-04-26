package main

import (
	"fmt"
	"log"

	"ZpcGoProject/config"
	"ZpcGoProject/router"
)

func main() {
	config.InitNacosConfig()
	config.InitDB()
	config.RegisterService()

	r := router.Init()
	addr := ":" + config.ServerPort
	fmt.Println("service started on " + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
