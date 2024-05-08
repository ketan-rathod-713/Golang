package category

import (
	"graphql_search/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database *mongo.Database
}

func New(db *mongo.Database) *api {
	return &api{
		Database: db,
	}
}

type Api interface {
	// define interface methods
	Create(name string) (*models.Category, error)
	Get(id string) (*models.Category, error)
	GetAll(pagination *models.Pagination) ([]*models.Category, error)
}
