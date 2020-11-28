package config

import "os"

type Config struct {
	HttpPort      string
	RedisAddress  string
	RedisPassword string
	AdminPassword string
}

func New() *Config {
	return &Config{
		HttpPort:      ParseEnv("HTTP_PORT", "8080"),
		RedisAddress:  ParseEnv("REDIS_ADDRESS", "localhost:6379"),
		RedisPassword: ParseEnv("REDIS_PASSWORD", ""),
		AdminPassword: ParseEnv("ADMIN_PASSWORD", "admin"),
	}
}

func ParseEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return def
}
