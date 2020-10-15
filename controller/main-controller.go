package controller

import "github.com/gofiber/fiber/v2"

// RouteMain -> prefix `/main`
func RouteMain(app *fiber.App) {
	mainGroup := app.Group("/")
	mainGroup.Get("/", helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}
