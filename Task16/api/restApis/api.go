package restapis

import (
	"graphql_search/models"
	"graphql_search/service/auth"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database       *mongo.Database
	DB_Collections *models.DB_COLLECTIONS
	AuthService    auth.Service
}

func New(database *mongo.Database, dbCollections *models.DB_COLLECTIONS, authService auth.Service) *api {
	return &api{
		Database:       database,
		DB_Collections: dbCollections,
		AuthService:    authService,
	}
}

type Api interface {
	VerifyEmail(w http.ResponseWriter, r *http.Request)
}
