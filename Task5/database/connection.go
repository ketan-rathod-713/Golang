package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Environment Variables Required
type DBConfig struct { // can i define it in main and then use it.
	DB_USER          string
	DB_USER_PASSWORD string
	DATABASE         string
	HOST             string
	PORT             string
	DB_PORT          string
}

func ConnectDb(config *DBConfig) *sql.DB {

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get environment variables
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_USER_PASSWORD = os.Getenv("DB_USER_PASSWORD")
	config.DATABASE = os.Getenv("DATABASE")
	config.HOST = os.Getenv("HOST")
	config.PORT = os.Getenv("PORT")
	config.DB_PORT = os.Getenv("DB_PORT")

	URL := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", config.DB_USER, config.DB_USER_PASSWORD, config.HOST, config.DATABASE)

	db, err := sql.Open("postgres", URL)
	CheckError(err)

	// defer db.Close()

	CheckError(db.Ping())

	return db
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
