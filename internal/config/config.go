package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port      int
	IdealTime int
	WriteTime int
	ReadTime  int
	Database  DBConfig
	Redis     RedisConfig
}
type DBConfig struct {
	DSN string
}
type RedisConfig struct {
	URL  string
	Port int
}

func getEnv(key string, required bool) string {
	res := os.Getenv(key)
	if res == "" && required {
		panic(fmt.Sprintf("requred key not found key: %s", key))
	}
	return res
}

func getEnvInt(key string, required bool) int {
	value := getEnv(key, required)
	res, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return res
}

func init() {
	if os.Getenv("APP_ENV") == "" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:      getEnvInt("port", true),
		IdealTime: getEnvInt("ideal_time", true),
		WriteTime: getEnvInt("write_time", true),
		ReadTime:  getEnvInt("read_time", true),
		Database: DBConfig{
			DSN: getEnv("db_dsn", true),
		},
		Redis: RedisConfig{
			URL:  getEnv("redis_url", true),
			Port: getEnvInt("redis_port", true),
		},
	}
}
