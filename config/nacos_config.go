package config

import (
	"log"

	"gopkg.in/yaml.v3"
)

func InitNacosConfig() {
	content, err := fetchNacosConfig("user-service.yaml", "DEFAULT_GROUP")
	if err != nil {
		log.Fatalf("配置获取失败: %v", err)
	}

	type AllConfig struct {
		Server   struct{ Port string }
		ApiToken string
		Mysql    MysqlConfig
		App      AppConfig
	}

	var cfg AllConfig
	if err := yaml.Unmarshal([]byte(content), &cfg); err != nil {
		log.Fatalf("配置解析失败: %v", err)
	}

	ServerPort = cfg.Server.Port
	ApiToken = cfg.ApiToken
	DBConfig = cfg.Mysql
	App = cfg.App
}
