package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nahkar/money-keeper/internal/db"
	"github.com/nahkar/money-keeper/internal/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	PORT := os.Getenv("PORT")

	dbConnection := db.Connect()
	defer dbConnection.Close()

	app := fiber.New()

	router.SetupRoutes(app, dbConnection)

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
