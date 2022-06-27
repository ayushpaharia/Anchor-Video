package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Something(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "Hello",
		"success": true,
	})
}
