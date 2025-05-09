package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type APPConfig struct {
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JwtSecret string
}

func LoadConfig() APPConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return APPConfig{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "user"),
		DBPass:    getEnv("DB_PASS", "password"),
		DBName:    getEnv("DB_NAME", "dbname"),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
