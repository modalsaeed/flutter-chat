package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	ServerPort string

	EmailDomain string
	EmailAPIKey string
	EmailFrom   string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using enviornment variables")
	}

	cfg := &Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		EmailDomain: os.Getenv("EMAIL_DOMAIN"),
		EmailAPIKey: os.Getenv("EMAIL_API_KEY"),
		EmailFrom:   os.Getenv("EMAIL_FROM"),
	}

	missing := false

	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		log.Println("Missing required database enviornment variables")
		missing = true
	}
	if cfg.ServerPort == "" {
		log.Println("Missing required server port")
		missing = true
	}
	if cfg.EmailDomain == "" || cfg.EmailAPIKey == "" || cfg.EmailFrom == "" {
		log.Println("Missing required email enviornment variables")
		missing = true
	}

	if missing {
		log.Fatal("One or more required enviornment variables are missing")
	}

	return cfg
}
