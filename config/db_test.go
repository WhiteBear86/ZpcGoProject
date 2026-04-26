package config

import "testing"

func TestMysqlConfigDSN(t *testing.T) {
	cfg := MysqlConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "secret",
		Database: "zpc",
		Charset:  "utf8mb4",
	}

	got := cfg.DSN()
	want := "root:secret@tcp(127.0.0.1:3306)/zpc?charset=utf8mb4&parseTime=True&loc=Local"
	if got != want {
		t.Fatalf("DSN() = %q, want %q", got, want)
	}
}
