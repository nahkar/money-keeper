package user

import (
	"errors"

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
		return utils.HandleError(c, err, fiber.StatusInternalServerError)

	}
	return c.JSON(users)
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}

	user, err := h.Service.GetUser(id)

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return utils.HandleError(c, err, fiber.StatusNotFound)
		}
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	return c.JSON(user)
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	var input CreateUserRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}

	if errors, err := utils.ValidateStruct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Validation failed",
			"fields": errors,
		})
	}

	user, err := h.Service.CreateUser(input)
	if err != nil {
		return utils.HandleError(c, utils.MapSQLError(err), fiber.StatusInternalServerError)
	}
	return c.JSON(user)
}

func (h *UserHandlers) UpdateUser(c *fiber.Ctx) error {
	var input UpdateUserRequest

	if err := c.BodyParser(&input); err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}

	if errors, err := utils.ValidateStruct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Validation failed",
			"fields": errors,
		})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}

	user, err := h.Service.UpdateUser(id, input)
	if err != nil {
		return utils.HandleError(c, utils.MapSQLError(err), fiber.StatusInternalServerError)
	}
	return c.JSON(user)
}

func (h *UserHandlers) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.HandleError(c, err, fiber.StatusBadRequest)
	}

	if err := h.Service.DeleteUser(id); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return utils.HandleError(c, err, fiber.StatusNotFound)
		}
		return utils.HandleError(c, err, fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
