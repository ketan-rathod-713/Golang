package userapi

import (
	"auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	DB     *mongo.Database
	Config *models.Config
}

func NewApi(db *mongo.Database, config *models.Config) *Api {

	// get jwt auth service

	return &Api{
		DB:     db,
		Config: config,
	}
}
