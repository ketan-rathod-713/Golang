package main

import (
	"graphql_search/api"
	"graphql_search/database"
	"graphql_search/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	configs, dataCollections, err := database.LoadEnv()

	// panics if unable to load environment variables
	if err != nil {
		log.Fatal(err.Error())
	}

	client, err := database.Connect(configs)

	if err != nil {
		log.Fatal(err.Error())
	}

	// static content
	fileServerHandler := http.FileServer(http.Dir("public"))
	http.Handle("/", fileServerHandler)

	resolver := &graph.Resolver{Api: api.New(client, configs, dataCollections)}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", configs.PORT)
	log.Printf("connect to http://localhost:%s/ for Home Page", configs.PORT)
	log.Fatal(http.ListenAndServe(":"+configs.PORT, nil))
}
