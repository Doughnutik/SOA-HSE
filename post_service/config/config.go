package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Printf("Ошибка чтения файла конфигурации: %v", err)
	}

	return &Config{
		Port: getEnv("PORT", "8082"),
		Host: getEnv("HOST", "127.0.0.1"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
