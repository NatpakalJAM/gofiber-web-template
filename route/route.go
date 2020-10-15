package route

import (
	"gofiber-web-template/controller"

	"github.com/gofiber/fiber/v2"
)

// Init -> init route
func Init(app *fiber.App) {
	mainGroup := app.Group("/")
	mainGroup.Get("/", controller.Index)
	mainGroup.Get("/hello_world", controller.HelloWorld)
}
