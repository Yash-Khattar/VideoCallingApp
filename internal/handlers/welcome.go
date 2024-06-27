package handlers

import(
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error{
	// renders html page under layout/main
	return c.Render("welcome", nil, "layout/main")
}