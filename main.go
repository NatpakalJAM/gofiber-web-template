package main

import (
	"fmt"
	"gofiber-web-template/cfg"
	"gofiber-web-template/route"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/utahta/go-cronowriter"
)

func init() {
	cfg.Init()
	//db.Init()
	//redis.Init()
	//queue.Init()
}

func main() {
	engine := html.New("./views", ".html")
	prefork := false // fiber config Prefork | Default: false
	if cfg.C.Environment != "development" {
		prefork = true
	} else {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Prefork:      prefork,
		Views:        engine,
		ErrorHandler: errorHandler(),
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: http.Dir("./public"),
	}))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(requestid.New())
	app.Use(logger.New(prepareLogger()))

	route.Init(app)

	// show error when route not found
	app.Use(func(c *fiber.Ctx) error {
		err := c.Status(fiber.StatusNotFound).SendFile(fmt.Sprintf("./views/error/%d.html", fiber.StatusNotFound))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("404 Not Found.")
		}
		return nil
	})

	// go func() {
	// 	queue.StartProcessConsumer()
	// }()

	log.Fatal(app.Listen(":3000"))
}

func prepareLogger() (config logger.Config) {
	outfile := cronowriter.MustNew("./logs/%Y-%m-%d.log", cronowriter.WithMutex())
	config.Format = "[${time}] ${status} - ${latency} ${method} ${path} query:[${query:}] header:[${header:}] body:[${body}]\n"
	config.Output = outfile
	return config
}

func errorHandler() func(*fiber.Ctx, error) error {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		err = ctx.Status(code).SendFile("./views/error/error.html")
		if err != nil {
			return ctx.Status(500).SendString("Internal Server Error")
		}
		return nil
	}
}
