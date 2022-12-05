package links

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"hungon.space/xurl/app/links/dto"
	"hungon.space/xurl/app/models"
	"hungon.space/xurl/common/utils"
	"hungon.space/xurl/packages/databases"
)

func CreateLink(c *fiber.Ctx) error {
	linkRepo := databases.NewLinkRepository()
	req := new(dto.CreateLinkDTO)
	apikey := c.GetReqHeaders()["apikey"]
	lastId, err := linkRepo.LastId()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	nextId := lastId + 1
	linkId := utils.LinkId(nextId)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	link := &models.Link{
		OriginalLink: req.OriginalLink,
		ApiKey:       apikey,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LinkId:       linkId,
	}

	link, err = linkRepo.CreateLink(link)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"link": link,
	})
}
