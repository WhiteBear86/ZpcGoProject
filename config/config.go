package config

import "fmt"

var (
	ServerPort string
	ApiToken   string
	DBConfig   MysqlConfig
	App        AppConfig
)

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}

func (c MysqlConfig) DSN() string {
	charset := c.Charset
	if charset == "" {
		charset = "utf8mb4"
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		charset,
	)
}

type AppConfig struct {
	Name string
	Env  string
}
