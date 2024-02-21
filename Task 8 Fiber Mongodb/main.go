package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApiStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func homeHandler(ctx *fiber.Ctx) error {
	// default cType is application/json for this
	return ctx.JSON(ApiStatus{Status: "OK", Message: "Api Working Fine"})
}

func main() {
	// * Connect to mongodb

	// TODO what is use case of this
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://Ketan:Ketan17547@@localhost:27017"))
	if err != nil {
		log.Error(err)
	}

	// *mongo.Collection // *mongo.Database
	collection := client.Database("bacancy").Collection("users")
	app := fiber.New()

	// A Collection can be used to query the database or insert documents:
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		log.Error(err)
	}

	id := res.InsertedID

	log.Info("Id inserted is ", id)

	app.Get("/", homeHandler)

	// Listen On Port Given In .env file
	app.Listen(":8080")
}
