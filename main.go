package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/packages/controller"
	"hungon.space/xurl/packages/middleware"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		Prefork: os.Getenv("PREFORK") == "true",
	})

	// Middlewares
	middleware.AppMiddlewares(app)

	// Routes
	controller.AppController(app)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
