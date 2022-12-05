package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/app/links"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	// ----- LINKS -----

	// Create a new link
	route.Post("/links", func(c *fiber.Ctx) error {
		requestId := c.GetReqHeaders()["requestId"]
		fmt.Println(requestId)
		return links.CreateLink(c)
	})

	// Get original link by short link
	route.Get("/:linkId", func(c *fiber.Ctx) error {
		return links.GetOriginalLink(c)
	})
}
