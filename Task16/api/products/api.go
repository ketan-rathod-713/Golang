package products

import (
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
	// define interface methods
	Create(name string, description string, price float64, quantity int, category string) (*models.Product, error)
	GetAll(pagination *models.Pagination) ([]*models.Product, error)
	Get(id string) (*models.Product, error)
}
