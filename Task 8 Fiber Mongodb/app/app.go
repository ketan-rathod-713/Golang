package app

import (
	"fibermongoapp/configs"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB *mongo.Client
}

/* Initialize App (configuring database, load envs and all ) */
func New() (*App, error) {
	// Load environment variables
	err := configs.LoadEnvFile()
	if err != nil {
		log.Error("Failed to load env variables")
		return nil, err
	}

	db, err := configs.ConnectDB()
	if err != nil {
		log.Error("Error Connecting Database")
		return nil, err
	}

	return &App{DB: db}, nil
}
