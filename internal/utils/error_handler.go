package utils

import (
	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": err.Error()})
}
