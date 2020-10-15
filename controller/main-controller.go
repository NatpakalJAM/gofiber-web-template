package controller

import "github.com/gofiber/fiber/v2"

// HelloWorld page
func HelloWorld(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}
