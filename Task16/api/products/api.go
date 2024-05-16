package products

import (
	"context"
	"graphql_search/models"
	"graphql_search/service/cashe"
	"graphql_search/service/cashe/product"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database            *mongo.Database
	DB_Collections      *models.DB_COLLECTIONS
	RedisProductService product.Service
}

func New(database *mongo.Database, dbCollections *models.DB_COLLECTIONS, redisService *cashe.Service) *api {
	return &api{
		Database:            database,
		DB_Collections:      dbCollections,
		RedisProductService: redisService.Product,
	}
}

type Api interface {
	// define interface methods
	Create(ctx context.Context, name string, description string, price float64, quantity int, category string) (*models.Product, error)

	// Authorizd request // normal user can see only sell products // admins can see all including archived, and new products too.
	GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.Product, error)
	Get(ctx context.Context, id string) (*models.Product, error)

	// UpdateProduct() // update qty, name, price, status,
}
