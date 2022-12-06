package links

import (
	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/common/logger"
)

func LinkController(r fiber.Router) {
	r.Post("/links", func(c *fiber.Ctx) error {
		logger.Info(c, "SHORTEN_LINK")
		return ShortenLink(c)
	})
}
