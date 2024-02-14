package main

import (
	"log"
	"task6MuxGorm/api"
	"task6MuxGorm/app"
)

func main() {
	// TODO:Initialise App with DB
	appInstance, err := app.NewApp()
	if err != nil {
		log.Println("ERROR: unable to create app instance")
	}

	// TODO: Intialise API With Initialise All Routes and services abstraction
	myApi, err := api.NewApi(appInstance)
	if err != nil {
		log.Println("ERROR: unable to initialize API")
	}

	log.Println(myApi)
	// TODO: Start Server On given router
}
