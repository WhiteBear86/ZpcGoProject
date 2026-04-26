package config

import "testing"

func TestParseNacosConfigReadsCamelCaseAPIToken(t *testing.T) {
	content := `
server:
  port: "8080"
apiToken: "secret-token"
mysql:
  host: "127.0.0.1"
  port: "3306"
  username: "root"
  password: "pwd"
  database: "demo"
app:
  name: "user-service"
  env: "dev"
`

	cfg, err := parseNacosConfig(content)
	if err != nil {
		t.Fatalf("parseNacosConfig returned error: %v", err)
	}

	if cfg.ApiToken != "secret-token" {
		t.Fatalf("ApiToken = %q, want %q", cfg.ApiToken, "secret-token")
	}
}
