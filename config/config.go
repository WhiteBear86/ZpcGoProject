package config

import "fmt"

/*
定义变量
*/
var (
	ServerPort string //服务端口
	ApiToken   string //apiToken
	DBConfig   MysqlConfig
	App        AppConfig
)

/*
Mysql连接信息的结构体
*/
type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}

/*
*
组装Mysql连接url的方法
*/
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

/*
*
应用信息的结构体
*/
type AppConfig struct {
	Name string
	Env  string
}
