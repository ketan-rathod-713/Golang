package database

import (
	"context"
	"errors"
	"graphql_search/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *models.Configs) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DB_URL))

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Load environment variables
func LoadEnv() (*models.Configs, *models.DB_COLLECTIONS, error) {
	godotenv.Load()

	// Check if particular values are there in .env file, if not then generate an error.

	configs := &models.Configs{
		PORT:     os.Getenv("PORT"),
		DB_URL:   os.Getenv("DB_URL"),
		DATABASE: os.Getenv("DATABASE"),
	}

	dbCollections := &models.DB_COLLECTIONS{
		CATEGORY: os.Getenv("COLL_CATEGORY"),
		PRODUCTS: os.Getenv("COLL_PRODUCTS"),
		CART:     os.Getenv("COLL_CARTS"),
		ORDERS:   os.Getenv("COLL_ORDERS"),
		USERS: os.Getenv("COLL_USERS"),
		LOGS: os.Getenv("COLL_LOGS"),
	}

	if configs.PORT == "" || configs.DB_URL == "" {
		return nil, nil, errors.New(".env file is not upto mark.")
	}

	return configs, dbCollections, nil
}
