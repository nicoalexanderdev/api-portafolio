package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	URI     string
	Name    string
	Timeout time.Duration
}

var (
	config *Configuration
	once   sync.Once
)

func GetConfig() *Configuration {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
		config = &Configuration{
			Server: ServerConfig{
				Port: getEnv("SERVER_PORT", "8080"),
				Mode: getEnv("GIN_MODE", "debug"),
			},
			Database: DatabaseConfig{
				URI:     getEnv("MONGODB_URI", ""),
				Name:    getEnv("MONGODB_DATABASE", "api_go"),
				Timeout: time.Duration(getEnvAsInt("MONGODB_TIMEOUT", 10)) * time.Second,
			},
		}
	})
	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
