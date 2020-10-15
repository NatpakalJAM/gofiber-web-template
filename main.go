package main

import (
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
		engine.Reload(true)
		engine.Debug(true)
		prefork = true
	}

	app := fiber.New(fiber.Config{
		Prefork: prefork,
		Views:   engine,
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