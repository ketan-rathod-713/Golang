package main

import (
	"log"
	"meetmeup/graph"
	"meetmeup/postgres"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8081"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "bacancy",
		Password: "admin",
		Database: "graphqlexample",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configs := graph.Config{Resolvers: &graph.Resolver{
		MeetupRepo: postgres.MeetupRepo{DB: DB},
		UserRepo:   postgres.UserRepo{DB: DB},
	}}


	srv := handler.NewDefaultServer(graph.NewExecutableSchema(configs))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	
	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
