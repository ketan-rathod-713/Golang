package user

import (
	"context"
	"graphql_search/models"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database       *mongo.Database
	DB_Collections *models.DB_COLLECTIONS
	Validator      *validator.Validate
}

func New(database *mongo.Database, dbCollections *models.DB_COLLECTIONS) *api {
	return &api{
		Database:       database,
		DB_Collections: dbCollections,
		Validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

type Api interface {
	RegisterUser(ctx context.Context, name string, emailID string, phoneNumber string, address models.AddressInput) (*models.User, error)
	SignInUser(ctx context.Context, id string) (*models.User, error)
	GetAllUsers(ctx context.Context, authToken string) ([]*models.User, error)
}
