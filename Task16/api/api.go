package api

import (
	"graphql_search/api/board"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	Client   *mongo.Client
	BoardApi board.BoardApi
}

func New(client *mongo.Client) *Api {
	return &Api{
		Client:   client,
		BoardApi: board.New(client.Database("task16")),
	}
}
