package config

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type nacosConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	ApiToken string      `yaml:"apiToken"`
	Mysql    MysqlConfig `yaml:"mysql"`
	App      AppConfig   `yaml:"app"`
}

func parseNacosConfig(content string) (nacosConfig, error) {
	var cfg nacosConfig
	if err := yaml.Unmarshal([]byte(content), &cfg); err != nil {
		return nacosConfig{}, fmt.Errorf("parse nacos config: %w", err)
	}
	return cfg, nil
}

func InitNacosConfig() {
	content, err := fetchNacosConfig("user-service.yaml", "DEFAULT_GROUP")
	if err != nil {
		log.Fatalf("閰嶇疆鑾峰彇澶辫触: %v", err)
	} else {
		log.Println("閰嶇疆鑾峰彇鎴愬姛")
	}

	cfg, err := parseNacosConfig(content)
	if err != nil {
		log.Fatalf("閰嶇疆瑙ｆ瀽澶辫触: %v", err)
	}

	ServerPort = cfg.Server.Port
	ApiToken = cfg.ApiToken
	DBConfig = cfg.Mysql
	App = cfg.App
}
