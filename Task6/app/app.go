package app

import (
	"task6MuxGorm/database"

	"gorm.io/gorm"
)

// App will hold DB connection which can be used in whole app
type App struct {
	DB *gorm.DB
}

// Connect To Database and return App object
func NewApp() (app *App, err error) {
	var db *gorm.DB
	db, err = database.InitialiseDB()
	if err != nil {
		return nil, err
	}

	app = &App{DB: db}

	return app, nil
}
