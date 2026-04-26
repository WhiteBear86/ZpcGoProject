package config

import "log"

func RegisterService() {
	port, err := parseServicePort(ServerPort)
	if err != nil {
		log.Printf("服务端口无效，跳过 Nacos 注册: %v", err)
		return
	}

	if err := registerServiceHTTP(App.Name, "127.0.0.1", port, "DEFAULT_GROUP"); err != nil {
		log.Printf("服务注册失败: %v", err)
		return
	}

	log.Printf("服务已注册到 Nacos: %s:%s", App.Name, ServerPort)
}
