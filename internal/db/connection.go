package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	host := os.Getenv("MONEY_KEEPER_DB_HOST")
	port := os.Getenv("MONEY_KEEPER_DB_PORT")
	user := os.Getenv("MONEY_KEEPER_DB_USER")
	password := os.Getenv("MONEY_KEEPER_DB_PASSWORD")
	dbname := os.Getenv("MONEY_KEEPER_DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	log.Println("Connected to the database successfully!")
	return db
}
