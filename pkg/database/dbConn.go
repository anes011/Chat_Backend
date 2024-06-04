package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbString := os.Getenv("DB_CONN")

	db, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	return db
}
