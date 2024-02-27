package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func WithAuthenticatedUser(c *fiber.Ctx) error {
	fmt.Printf("Stuff ltuff\n\n")
	return c.Next()
}
