package board

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

type BoardApi interface {
	GetBoardsByTitle(title string) ([]*models.Board, error)
}
