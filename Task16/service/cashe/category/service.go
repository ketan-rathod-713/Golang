package category

import (
	"graphql_search/models"

	"github.com/go-redis/redis/v8"
)

// what type of interface should i provide for redis ??
type service struct {
	Client *redis.Client
}

func New(client *redis.Client) Service {
	return &service{
		Client: client,
	}
}

// This is the contract for the api's using this service.
type Service interface {
	Get(key string) (*models.CategoryDB, error)
	Set(productDb *models.CategoryDB) error
}
