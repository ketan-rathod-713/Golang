package app

import (
	authGrpc "proto/auth-server/v1"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB         *mongo.Database
	AuthClient authGrpc.AuthClient
}

func New(db *mongo.Database, client authGrpc.AuthClient) *App {
	return &App{
		DB:         db,
		AuthClient: client,
	}
}
