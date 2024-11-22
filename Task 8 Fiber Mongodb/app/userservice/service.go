package userservice

// ! Collection name is hardcoded here.

import (
	"fibermongoapp/configs"
	"fibermongoapp/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service interface will serve as a contract for user services
type Service interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUsers(queries map[string]string) ([]*models.User, error)
	GetOneUserById(userId primitive.ObjectID) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(userId primitive.ObjectID) (*models.User, error)
}

/*
user service requires db pointer and mongo.Collection pointer to query user collection
TODO: here collection name is hardcoded -> improve it.
*/
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
