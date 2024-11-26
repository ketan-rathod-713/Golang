package main

import (
	"fmt"
	"log"
	"net/http"
	"task6MuxGorm/api"
	"task6MuxGorm/app"
	_ "task6MuxGorm/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

// @summary Show a hello message
// @description get hello message
// @id hello
// @produce json
// @success 200 {string} string "ok"
// @router /hello [get]

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

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	mainApi.InitialiseRoutes(router)

	log.Println("Router Initialised", mainApp.Config.PORT)

	// TODO: Start Server On given router
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", mainApp.Config.PORT), router))

	// test.TestIssueBook()
}
