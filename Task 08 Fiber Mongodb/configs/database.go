package configs

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(ENV_MONGO_URI()))

	if err != nil {
		log.Info("ERROR connecting mongodb")
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		log.Info("ERROR Ping Mongodb")
		return nil, err
	}

	log.Info("Connected to Mongodb")
	return mongoClient, nil
}

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database(ENV_DATBASE()).Collection(collectionName)
    return collection
}