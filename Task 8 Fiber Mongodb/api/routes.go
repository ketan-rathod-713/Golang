package api

import (
	"fibermongoapp/api/adminapi"
	"fibermongoapp/api/userapi"
	"fibermongoapp/app"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/* api struct for wrapper around all api's */
type api struct {
	App *app.App
}

func New(a *app.App) *api {
	return &api{
		App: a,
	}
}

/* Initialize routes */
func (a *api) InitializeRoutes(app *fiber.App) {
	// intialize all routes
	app.Get("/", homeHandler) // ! if using group here then for any route this handler will be called

	/* user routes */
	userRouter := app.Group("/user")
	userApi := userapi.New(a.App)
	userApi.Routes(userRouter)

	/* other routes */
	adminRouter := app.Group("/admin")
	adminApi := adminapi.New(a.App)
	adminApi.Routes(adminRouter)

}

func homeHandler(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"Status": 200, "Message": "Server Running Fine"})
}
