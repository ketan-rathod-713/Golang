package main

import (
	"books/api"
	"books/app"
	"books/db"
	"books/grpcclient"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// load environment variables
	configs := db.LoadEnv()

	client, err := db.ConnectDB(configs)
	if err != nil {
		log.Fatal("Error Occured")
	}

	// TODO change this method of communication.
	// connect to grpc server
	conn, grpcClient, err := grpcclient.NewAuthGrpcClient("localhost:8081")
	defer conn.Close()

	if err != nil {
		log.Fatal("error connecting grpc client")
	}

	// initialise app
	App := app.New(client.Database(configs.DATABASE), grpcClient)

	// initialise api and routes
	Api := api.NewApi(App, configs)
	r := mux.NewRouter()
	Api.InitializeRoutes(r)

	// start REST and grpc server
	go func() {
		fmt.Println("Book server started on port", configs.REST_PORT)

		err = http.ListenAndServe(configs.REST_PORT, r)

		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		fmt.Println("Currently No GRPC serices are available")

	}()

	select {}
}

// use http only
