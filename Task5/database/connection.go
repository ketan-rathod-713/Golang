package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env variable
	DB_USER := os.Getenv("DB_USER")
	DB_USER_PASSWORD := os.Getenv("DB_USER_PASSWORD")
	DATABASE := os.Getenv("DATABASE")

	URL := fmt.Sprintf("postgres://%v:%v@localhost/%v?sslmode=disable", DB_USER, DB_USER_PASSWORD, DATABASE)

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
