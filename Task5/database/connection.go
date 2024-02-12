package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const URL = "postgres://bacancy:admin@localhost/bacancy?sslmode=disable"

func ConnectDb() *sql.DB {
	db, err := sql.Open("postgres", URL)
	CheckError(err)

	// defer db.Close()

	CheckError(db.Ping())

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
