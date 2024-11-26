package utils

import (
	"os"
	"wikipediasearch/models"

	"github.com/joho/godotenv"
)

func LoadEnv() *models.Config {
	godotenv.Load()

	return &models.Config{
		PORT:     os.Getenv("PORT"),
		WIKI_URL: os.Getenv("WIKI_URL"),
	}
}
