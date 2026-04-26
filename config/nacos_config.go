package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
	"log"
)

func InitNacosConfig() {
	sc := []constant.ServerConfig{{IpAddr: "127.0.0.1", Port: 8800}}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		Username:            "nacos",
		Password:            "nacos",
	}

	client, err := clients.NewConfigClient(vo.NacosClientParam{ClientConfig: &cc, ServerConfigs: sc})
	if err != nil {
		log.Fatalf("Nacos连接失败: %v", err)
	}

	content, err := client.GetConfig(vo.ConfigParam{DataId: "user-service.yaml", Group: "DEFAULT_GROUP"})
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
	_ = yaml.Unmarshal([]byte(content), &cfg)

	ServerPort = cfg.Server.Port
	ApiToken = cfg.ApiToken
	DB = cfg.Mysql
	App = cfg.App
}
