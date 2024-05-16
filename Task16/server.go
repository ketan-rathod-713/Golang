package main

import (
	"graphql_search/api"
	"graphql_search/database"
	"graphql_search/graph"
	"graphql_search/service/cashe"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	configs, dataCollections, err := database.LoadEnv()
	// panics if unable to load environment variables
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Environment variables loaded.")

	client, err := database.Connect(configs)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mongodb database connected.")

	redisClient, err := cashe.Init("", "", 0, "localhost:6379")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Redis database connected.")

	mx := mux.NewRouter()

	api := api.New(client, configs, dataCollections, redisClient)
	api.InitializeRoutes(mx)

	// static content
	fileServerHandler := http.FileServer(http.Dir("public"))
	mx.Handle("/", fileServerHandler)

	resolver := &graph.Resolver{Api: api}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	mx.Use(api.CategoryApi.CategoryLoaderMiddleware)

	mx.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mx.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", configs.PORT)
	log.Printf("connect to http://localhost:%s/ for Home Page", configs.PORT)
	log.Fatal(http.ListenAndServe(":"+configs.PORT, mx))
}
