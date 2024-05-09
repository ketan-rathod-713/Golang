package products

import (
	"graphql_search/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database *mongo.Database
	DB_Co
}

func New(database *mongo.Database) *api {
	return &api{
		Database: database,
	}
}

type Api interface {
	// define interface methods
	Create(name string, description string, price float64, quantity int, category string) (*models.Product, error)
	GetAll(pagination *models.Pagination) ([]*models.Product, error)
	Get(id string) (*models.Product, error)
}
