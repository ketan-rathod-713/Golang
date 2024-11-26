package app

import (
	"fibermongoapp/configs"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	User                *mongo.Collection
	Subject             *mongo.Collection
	Class               *mongo.Collection
	StudentRegistration *mongo.Collection
	Teacher             *mongo.Collection
}

type App struct {
	DB          *mongo.Client
	Collections *Collections
}

/* Initialize App (configuring database, load envs and all ) */
func New() (*App, error) {
	// Load environment variables
	err := configs.LoadEnvFile()
	if err != nil {
		log.Error("Failed to load env variables")
		return nil, err
	}

	db, err := configs.ConnectDB()
	if err != nil {
		log.Error("Error Connecting Database")
		return nil, err
	}

	return &App{DB: db, Collections: &Collections{
		User:                configs.GetCollection(db, "users"),
		Subject:             configs.GetCollection(db, "subjects"),
		Class:               configs.GetCollection(db, "classes"),
		StudentRegistration: configs.GetCollection(db, "student_registrations"),
		Teacher:             configs.GetCollection(db, "teacher"),
	}}, nil
}
