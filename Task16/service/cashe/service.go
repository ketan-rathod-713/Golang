package cashe

import (
	"context"
	"graphql_search/service/cashe/category"
	"graphql_search/service/cashe/product"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	Client   *redis.Client
	Category category.Service
	Product  product.Service
}

func Init(username string, password string, database int, address string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

func New(client *redis.Client) *Service {
	return &Service{
		Client:   client,
		Category: category.New(client),
		Product:  product.New(client),
	}
}
