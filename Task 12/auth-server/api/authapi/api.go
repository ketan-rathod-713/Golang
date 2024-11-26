package authapi

import (
	"auth/app/jwtauth"
	"auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	DB         *mongo.Database
	Config     *models.Config
	JwtService jwtauth.Service
}

func NewApi(db *mongo.Database, config *models.Config) *Api {

	// get jwt auth service

	return &Api{
		DB:         db,
		Config:     config,
		JwtService: jwtauth.New(),
	}
}
