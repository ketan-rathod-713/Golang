package models

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT        string
	HOST        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
}

func LoadEnv() (*Env, error) {
	// load all environement variables
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Env{
		PORT:        os.Getenv("PORT"),
		HOST:        os.Getenv("HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
	}, nil
}
