package userservice

// ! Collection name is hardcoded here.

import (
	"fibermongoapp/configs"
	"fibermongoapp/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetOneUserById(userId primitive.ObjectID) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(userId primitive.ObjectID) (*models.User, error)
}

type service struct {
	DB             *mongo.Client
	UserCollection *mongo.Collection
}

func New(db *mongo.Client) Service {
	return &service{
		DB:             db,
		UserCollection: configs.GetCollection(db, "users"),
	}
}
