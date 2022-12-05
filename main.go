package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/packages/middlewares"
	"hungon.space/xurl/packages/routes"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		Prefork: os.Getenv("PREFORK") == "true",
	})

	// Middlewares
	middlewares.ApplyRequestIdMiddleware(app)

	// Routes
	routes.PublicRoutes(app)
	app.Listen(":" + os.Getenv("APP_PORT"))
}
