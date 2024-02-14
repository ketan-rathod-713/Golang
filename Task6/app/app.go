package app

import (
	"fmt"
	"log"
	"os"
	"task6MuxGorm/database"
	"task6MuxGorm/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// App will hold DB connection which can be used in whole app
type App struct {
	DB     *gorm.DB
	Config *models.Config
}

// Connect To Database and return App object
func NewApp() (app *App, err error) {
	var db *gorm.DB
	db, err = database.InitialiseDB()

	if err != nil {
		return nil, err
	}

	app = &App{DB: db, Config: loadEnv()}

	// After Getting Config And DB initialise and automigrate some stuff
	schemaQuery := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %v", app.Config.DB_SCHEMA_NAME)
	result := db.Exec(schemaQuery)

	if result.Error != nil {
		return app, result.Error
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return app, err
	}

	return app, nil
}

func loadEnv() *models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: unable to load.env file")
	}

	config := &models.Config{
		DB_PORT:          os.Getenv("DB_PORT"),
		DATABASE:         os.Getenv("DATABASE"),
		HOST:             os.Getenv("HOST"),
		DB_USER:          os.Getenv("DB_USER"),
		DB_USER_PASSWORD: os.Getenv("DB_USER_PASSWORD"),
		DB_SCHEMA_NAME:   os.Getenv("DB_SCHEMA_NAME"),
		PORT:             os.Getenv("PORT"),
	}

	return config
}
