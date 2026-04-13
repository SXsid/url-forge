package config

import "time"

type ServerConfig struct {
	Port      int
	IdealTime time.Duration
	WriteTime time.Duration
	ReadTime  time.Duration
	Database  DBConfig
	Redis     DBConfig
}
type DBConfig struct {
	DSN string
}
type RedisConfig struct {
	URL  string
	Port int
}

func NewServerConfig() (*ServerConfig, error) {
	return &ServerConfig{}, nil
}
