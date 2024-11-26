package main

import (
	"fmt"
	"log"
	"net/http"
	"task6MuxGorm/api"
	"task6MuxGorm/app"

	"github.com/gorilla/mux"
)

func main() {
	// TODO:Initialise App with DB, initialise environement variables
	mainApp, err := app.NewApp()
	if err != nil {
		log.Println("ERROR: unable to create app instance")
		log.Println("data got", mainApp.Config)
	}

	// TODO: Intialise API With Initialise All Routes and services abstraction
	mainApi, err := api.NewApi(mainApp)
	if err != nil {
		log.Println("ERROR: unable to initialize API")
	}

	log.Println("API Initialised")

	var router *mux.Router = mux.NewRouter()
	mainApi.InitialiseRoutes(router)

	log.Println("Router Initialised", mainApp.Config.PORT)

	// TODO: Start Server On given router
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", mainApp.Config.PORT), router))

	// test.TestIssueBook()
}
