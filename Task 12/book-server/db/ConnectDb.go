package db

import (
	"books/models"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(config *models.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DB_URL))

	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected To Mongodb")
	return mongoClient, nil
}

func LoadEnv() *models.Config {
	// Load All environment variables from the .env file in current directory
	godotenv.Load()

	// now make config
	config := models.Config{
		DB_URL:    os.Getenv("DB_URL"),
		REST_PORT: os.Getenv("REST_PORT"),
		GRPC_PORT: os.Getenv("GRPC_PORT"),
		SECRET:    os.Getenv("SECRET"),
		DATABASE:  os.Getenv("DATABASE"),
	}

	return &config
}
