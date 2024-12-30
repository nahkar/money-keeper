package user

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(db *sql.DB, routes fiber.Router) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandlers := NewUserHandlers(userService)

	routes.Get("/users", userHandlers.GetAllUsers)
	routes.Post("/users", userHandlers.CreateUser)
}
