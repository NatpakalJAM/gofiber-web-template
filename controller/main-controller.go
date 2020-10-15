package controller

import "github.com/gofiber/fiber/v2"

// Index page
func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Index Page",
	})
}

// HelloWorld page
func HelloWorld(c *fiber.Ctx) error {
	return c.Render("hello_world", fiber.Map{
		"Title": "Hello, World!",
	}, "layouts/main")
}
