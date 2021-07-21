package config

import (
	"Pigeon/library/env"
)

var MysqlDefaultConfig = &MysqlConfig{
	db:       env.GetOrDefault("DB_NAME", ""),
	port:     env.GetOrDefault("DB_PORT", ""),
	user:     env.GetOrDefault("DB_USER", ""),
	password: env.GetOrDefault("DB_PASSWORD", ""),
	charset:  "utf8mb4",
}

type MysqlConfig struct {
	db       string
	port     string
	user     string
	password string
	charset  string
}

func (c MysqlConfig) ToConfigString() string {
	return c.user + ":" + c.password + "@/" + c.db + "?charset=" + c.charset + "&parseTime=True&loc=Local"
}
