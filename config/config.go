package config

import (
	"log"
	"os"

	models "binance-integration/models"

	"github.com/joho/godotenv"
)

func LoadEnv() (cfg models.Configs) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg.MongoURI = os.Getenv("mongoURI")

	return cfg
}