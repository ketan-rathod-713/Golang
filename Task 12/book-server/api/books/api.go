package books

import (
	"books/app"
	"books/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

type Api struct {
	App    *app.App
	Config *models.Config
}

func NewApi(app *app.App, config *models.Config) *Api {
	return &Api{
		App:    app,
		Config: config,
	}
}
