package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func WithAuthenticatedUser(c *fiber.Ctx) error {
	log.Default().Println("Authenticated user")
	return c.Next()
}
