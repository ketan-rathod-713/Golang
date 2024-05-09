package user

import (
	"context"
	"graphql_search/graph/model"
	"graphql_search/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database       *mongo.Database
	DB_Collections *models.DB_COLLECTIONS
}

func New(database *mongo.Database, dbCollections *models.DB_COLLECTIONS) *api {
	return &api{
		Database:       database,
		DB_Collections: dbCollections,
	}
}

type Api interface {
	RegisterUser(ctx context.Context, name string, emailID string, phoneNumber string, address model.AddressInput) (*model.User, error)
	SignInUser(ctx context.Context, id string) (*model.User, error)
	GetAllUsers(ctx context.Context, authToken string) ([]*model.User, error)
}