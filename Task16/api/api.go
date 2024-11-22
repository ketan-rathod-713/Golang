package api

import (
	"graphql_search/api/board"
	"graphql_search/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	Client   *mongo.Client
	BoardApi board.BoardApi
}

func New(client *mongo.Client, configs *models.Configs) *Api {
	return &Api{
		Client:   client,
		BoardApi: board.New(client.Database(configs.DATABASE)),
	}
}
