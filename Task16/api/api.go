package api

import (
	"graphql_search/api/category"
	"graphql_search/api/products"
	restapis "graphql_search/api/restApis"
	"graphql_search/api/user"
	"graphql_search/models"
	"graphql_search/service/auth"
	"graphql_search/service/cashe"
	"graphql_search/service/mail"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	Client       *mongo.Client
	Configs      *models.Configs
	ProductApi   products.Api
	CategoryApi  category.Api
	UserApi      user.Api
	RedisService *cashe.Service
	RestApis     restapis.Api
}

func New(client *mongo.Client, configs *models.Configs, dbCollections *models.DB_COLLECTIONS, redisClient *redis.Client) *Api {

	var db = client.Database(configs.DATABASE)
	var redisService = cashe.New(redisClient)
	var authService = auth.New(configs.JWT_SECRET, 20*time.Minute)
	var mailService = mail.New(configs.SMTP_PASSWORD, configs.SMTP_EMAIL)

	return &Api{
		Client:      client,
		Configs:     configs,
		CategoryApi: category.New(db, dbCollections, redisService),
		ProductApi:  products.New(db, dbCollections, redisService),
		UserApi:     user.New(db, dbCollections, authService, mailService),
		RestApis:    restapis.New(db, dbCollections, authService),
	}
}

func (a *Api) InitializeRoutes(mx *mux.Router) {

	// pass authToken as query parameter.
	mx.HandleFunc("/verifyEmail", a.RestApis.VerifyEmail)
}
