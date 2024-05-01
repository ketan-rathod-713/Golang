package main

import (
	"log"
	"meetmeup/graph"
	"meetmeup/postgres"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

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

	var SandboxHTML = []byte(`
		<!DOCTYPE html>
		<html lang="en">
		<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
		<div id="sandbox" style="height:100vh; width:100vw;"></div>
		<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
		<script>
		new window.EmbeddedSandbox({
		target: "#sandbox",
		initialEndpoint: "http://localhost:8080/query",
		});
		</script>
		</body>
		</html>`)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(configs))

	// start file server
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	// For testing in local
	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(SandboxHTML)
	}))

	log.Printf("connect to http://localhost:%s/sandbox for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
