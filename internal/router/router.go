package router

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nahkar/money-keeper/internal/modules/user"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	routes := app.Group("/api")

	user.UserRoutes(db, routes)

}
