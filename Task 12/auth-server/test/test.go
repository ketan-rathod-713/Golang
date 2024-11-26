package test

import (
	"auth/db"
	"auth/models"
	"auth/api"
	"fmt"
)

func InitialiseTestEnvironment() (*api.Api, error) {
	// load environment variables
	// configs := db.LoadEnv()
	var configs = &models.Config{
		DB_URL:    "mongodb://localhost:27017/",
		REST_PORT: ":8080",
		GRPC_PORT: ":8081",
		SECRET:    "secret",
		DATABASE:  "grpcProjectTest",
	}

	fmt.Println("Configs got from env files are ", configs)

	client, err := db.ConnectDB(configs)
	if err != nil {
		return nil, err
	}

	Api := api.NewApi(client.Database(configs.DATABASE), configs)

	return Api, nil
}
