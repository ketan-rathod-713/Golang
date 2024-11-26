package app

import (
	"auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type app struct {
	DB     *mongo.Database
	Config *models.Config
}

func New(db *mongo.Database, config *models.Config) *app {
	return &app{
		DB:     db,
		Config: config,
	}
}
