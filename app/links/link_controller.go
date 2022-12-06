package links

import (
	"github.com/gofiber/fiber/v2"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
)

func LinkController(r fiber.Router) {

	linkService := LinkService{}
	// Link routes

	// SHORTEN LINK
	r.Post("/links", func(c *fiber.Ctx) error {
		logger.Info(c, "SHORTEN_LINK")
		apikey := c.Get("apikey")
		var body Link
		err := c.BodyParser(&body)
		if err != nil {
			return xerror.LinkBodyInvalid()
		}

		return linkService.ShortenLink(c, body.OriginalLink, apikey, body.Password, body.Type)
	})

	// GET LINK TYPE A
	r.Get("/a/:linkId", func(c *fiber.Ctx) error {
		logger.Info(c, "GET_LINK_TYPE_A")
		linkId := c.Params("linkId")
		return linkService.GetLinkTypeA(c, linkId)
	})
}
