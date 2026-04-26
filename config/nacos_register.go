package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

func RegisterService() {
	client, _ := clients.NewNamingClient(vo.NacosClientParam{})

	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8080,
		ServiceName: App.Name,
		GroupName:   "DEFAULT_GROUP",
	})
	if err != nil || !success {
		log.Println("服务注册失败")
	} else {
		log.Println("✅ 服务注册Nacos成功")
	}
}
