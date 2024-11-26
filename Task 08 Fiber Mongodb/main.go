package main

import (
	"fibermongoapp/api"
	"fibermongoapp/app"
	"fibermongoapp/configs"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	/* Initialize App (load env, connecting to database and app variables setup )*/
	App, err := app.New()
	if err != nil {
		log.Fatal("Error Initializing App ", err)
	}

	/* set up api using App */
	api := api.New(App)

	// setup fiber server
	app := fiber.New()

	/*App level Middlewares */
	app.Use(logger.New())

	/* set up routes for fiber app*/
	api.InitializeRoutes(app)

	//? Listen App On Given Port
	err = app.Listen(fmt.Sprintf(":%v", configs.ENV_PORT()))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Server Started On Port", configs.ENV_PORT())
	}
}
