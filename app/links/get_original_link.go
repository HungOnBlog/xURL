package links

import (
	"github.com/gofiber/fiber/v2"
)

func GetOriginalLink(c *fiber.Ctx) error {
	// Get link id from url
	linkId := c.Params("linkId")

	if linkId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "link id is required",
		})
	}

	return c.JSON(fiber.Map{
		"linkId": linkId,
	})
}
