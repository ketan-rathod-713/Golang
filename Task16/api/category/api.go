package category

import (
	"context"
	"graphql_search/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database          *mongo.Database
	DB_Collections    *models.DB_COLLECTIONS
	CategoryLoaderKey string
}

func New(db *mongo.Database, dbCollections *models.DB_COLLECTIONS) *api {
	return &api{
		Database:          db,
		DB_Collections:    dbCollections,
		CategoryLoaderKey: "categoryLoader",
	}
}

type Api interface {
	// define interface methods
	Create(ctx context.Context, name string) (*models.Category, error)
	Get(ctx context.Context, id string) (*models.Category, error)
	GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.Category, error)
	GetProductsByCategory(ctx context.Context, category *models.Category) ([]*models.Product, error)

	// category loader middleware
	CategoryLoaderMiddleware(next http.Handler) http.Handler
	// fetch categoryLoder from request's context.
	GetCategoryLoader(ctx context.Context) *CategoryLoader
}
