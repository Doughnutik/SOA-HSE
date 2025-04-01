package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	Host           string
	UserServiceUrl string
	PostServiceUrl string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Printf("Ошибка чтения файла конфигурации: %v", err)
	}

	return &Config{
		Port:           getEnv("PORT", "8081"),
		Host:           getEnv("HOST", "127.0.0.1"),
		UserServiceUrl: getEnv("USER_SERVICE_URL", "http://user_service:8080"),
		PostServiceUrl: getEnv("POST_SERVICE_URL", "http://post_service:8282"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
