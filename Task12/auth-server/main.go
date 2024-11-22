package main

import (
	"auth/api"
	"auth/db"
	"auth/grpcservice"
	"fmt"
	"log"
	"net"
	"net/http"

	grpcService "proto/auth-server/v1"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	// load environment variables
	configs := db.LoadEnv()

	client, err := db.ConnectDB(configs)
	if err != nil {
		log.Fatal("Error Connecting To Database")
	}

	// initialise api and initalise routes using mux router
	Api := api.NewApi(client.Database(configs.DATABASE), configs)
	r := mux.NewRouter()
	Api.InitializeRoutes(r)

	// start rest and grpc server
	go func() {
		fmt.Println("Auth server started on port", configs.REST_PORT)

		err = http.ListenAndServe(configs.REST_PORT, r)

		if err != nil {
			log.Fatal(err)
		}

	}()

	go func() {
		fmt.Println("Auth GRPC SERVICE started on port", configs.GRPC_PORT)

		lis, err := net.Listen("tcp", configs.GRPC_PORT)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		grpcService.RegisterAuthServer(s, grpcservice.New(configs))

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	// wait for infinite time
	select {}
}

// Folder Structure

// Api For All Api's
// App can have services if we want.
