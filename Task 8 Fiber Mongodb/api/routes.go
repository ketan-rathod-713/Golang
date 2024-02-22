package api

import (
	"fibermongoapp/api/userapi"
	"fibermongoapp/app"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

/* api struct for wrapper around all api's */
type api struct {
	DB *mongo.Client
}

func New(a *app.App) *api {
	return &api{
		DB: a.DB,
	}
}

/* Initialize routes */
func (a *api) InitializeRoutes(app *fiber.App) {
	// intialize all routes
	app.Get("/", homeHandler) // ! if using group here then for any route this handler will be called

	/* user routes */
	userRouter := app.Group("/user")
	userApi := userapi.New(a.DB)
	userApi.Routes(userRouter)

	/* other routes */

}

func homeHandler(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"Status": 200, "Message": "Server Running Fine"})
}
