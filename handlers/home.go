package handlers

import "github.com/gofiber/fiber/v2"

func HandleGetHome(c *fiber.Ctx) error {
	return c.Render("partials/first", fiber.Map{
	})
}
