package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		Prefork: os.Getenv("PREFORK") == "true",
	})

	app.Listen(":" + os.Getenv("APP_PORT"))
}
