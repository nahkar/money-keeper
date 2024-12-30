package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	Service *UserService
}

func NewUserHandlers(service *UserService) *UserHandlers {
	return &UserHandlers{Service: service}
}

var validate = validator.New()

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

	if err := validate.Struct(input); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Validation failed",
			"fields": errors,
		})
	}

	user, err := h.Service.CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
