package config

import (
	"os"
)

type Config struct {
	

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Redis
	RedisHost string
	RedisPort string
}

func Load() *Config {
	cfg := &Config{
		DBHost:    getEnv("POSTGRES_HOST", "db"),
		DBPort:    getEnv("POSTGRES_PORT", "5432"),
		DBUser:    getEnv("POSTGRES_USER", "postgres"),
		DBPassword:getEnv("POSTGRES_PASSWORD", "password"),
		DBName:    getEnv("POSTGRES_DB", "gokernel-db"),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
