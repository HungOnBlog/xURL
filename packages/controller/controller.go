package controller

import (
	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/app/links"
)

func AppController(a *fiber.App) {
	apiV1 := a.Group("/api/v1")

	// Link routes
	links.LinkController(apiV1)
}
