package api

import (
	"books/api/books"
	"books/app"
	"books/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

type api struct {
	App     *app.App
	Config  *models.Config
	BookApi *books.Api
}

func NewApi(app *app.App, config *models.Config) *api {
	return &api{
		App:     app,
		Config:  config,
		BookApi: books.NewApi(app, config),
	}
}
