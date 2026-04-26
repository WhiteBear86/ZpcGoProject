package config

var (
	ServerPort string
	ApiToken   string
	DB         MysqlConfig
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

type AppConfig struct {
	Name string
	Env  string
}
