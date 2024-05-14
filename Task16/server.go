package main

import (
	"graphql_search/api"
	"graphql_search/database"
	"graphql_search/graph"
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

	mx := mux.NewRouter()
	// static content
	fileServerHandler := http.FileServer(http.Dir("public"))
	mx.Handle("/", fileServerHandler)

	api := api.New(client, configs, dataCollections)

	resolver := &graph.Resolver{Api: api}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	mx.Use(api.CategoryApi.CategoryLoaderMiddleware)

	mx.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mx.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", configs.PORT)
	log.Printf("connect to http://localhost:%s/ for Home Page", configs.PORT)
	log.Fatal(http.ListenAndServe(":"+configs.PORT, mx))
}
