package api

import (
	"auth/api/authapi"
	"auth/api/userapi"
	"auth/app/jwtauth"
	"auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	DB         *mongo.Database
	Config     *models.Config
	JwtService jwtauth.Service
	AuthApi    *authapi.Api
	UserApi    *userapi.Api
}

func NewApi(db *mongo.Database, config *models.Config) *Api {

	// get jwt auth service

	return &Api{
		DB:         db,
		Config:     config,
		JwtService: jwtauth.New(),
		AuthApi:    authapi.NewApi(db, config),
		UserApi:    userapi.NewApi(db, config),
	}
}
