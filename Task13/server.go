package main

import (
	"fmt"
	"log"
	"meetmeup/graph"
	"meetmeup/models"
	"meetmeup/postgres"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	env, err := models.LoadEnv()

	if err != nil {
		log.Fatal("An Error Loading Environement Variables")
	}

	log.Println("Loaded Data From .ENV File")
	log.Println(env)

	DB := postgres.New(&pg.Options{
		User:     env.DB_USER,
		Password: env.DB_PASSWORD,
		Database: env.DB_DATABASE,
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	configs := graph.Config{Resolvers: &graph.Resolver{
		MeetupRepo: postgres.MeetupRepo{DB: DB},
		UserRepo:   postgres.UserRepo{DB: DB},
	}}

	var sandboxString = fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
		<div id="sandbox" style="height:100vh; width:100vw;"></div>
		<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
		<script>
		new window.EmbeddedSandbox({
		target: "#sandbox",
		initialEndpoint: "http://localhost:%v/query",
		});
		</script>
		</body>
		</html>`, env.PORT)

	var SandboxHTML = []byte(sandboxString)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(configs))

	// start file server
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)
	log.Printf("File Server Started On http://localhost:%v/", env.PORT)
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	// For testing in local
	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(SandboxHTML)
	}))

	log.Printf("connect to http://localhost:%s/sandbox for GraphQL playground", env.PORT)
	fmt.Println("")

	log.Fatal(http.ListenAndServe(":"+env.PORT, nil))

}
