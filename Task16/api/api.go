package api

import (
	"graphql_search/api/category"
	"graphql_search/api/products"
	"graphql_search/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	Client      *mongo.Client
	Configs     *models.Configs
	ProductApi  products.Api
	CategoryApi category.Api
}

func New(client *mongo.Client, configs *models.Configs) *Api {

	var db = client.Database(configs.DATABASE)

	return &Api{
		Client:  client,
		Configs: configs,
		CategoryApi: category.New(db),
		ProductApi: products.New(db),
	}
}
