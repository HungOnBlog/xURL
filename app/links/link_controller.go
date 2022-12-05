package links

import "github.com/gofiber/fiber/v2"

func LinkController(r fiber.Router) {
	r.Get("/links", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
