package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nahkar/money-keeper/internal/utils"
)

type UserHandlers struct {
	Service *UserService
}

func NewUserHandlers(service *UserService) *UserHandlers {
	return &UserHandlers{Service: service}
}

func (h *UserHandlers) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	var input CreateUserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if errors, err := utils.ValidateStruct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Validation failed",
			"fields": errors,
		})
	}

	user, err := h.Service.CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": utils.MapSQLError(err).Error()})
	}
	return c.JSON(user)
}
