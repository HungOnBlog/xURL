package links

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"hungon.space/xurl/common/logger"
	"hungon.space/xurl/common/utils"
)

func genHashId() (string, error) {
	lastId, err := LastLinkId()
	if err != nil {
		return "nil", err
	}

	return utils.HashId(lastId + 1), nil
}

func ShortenLink(c *fiber.Ctx) error {
	apikey := c.GetReqHeaders()["apikey"]
	linkId, error := genHashId()

	if error != nil {
		return error
	}

	link := new(Link)
	if err := c.BodyParser(link); err != nil {
		return err
	}

	link.LinkID = linkId
	link.ApiKey = apikey
	link.ShortLink = os.Getenv("BASE_URL") + "/" + linkId
	link.Type = "a"

	err := CreateLink(link)
	if err != nil {
		logger.Warn(c, "SHORTEN_LINK_FAILED", zap.Error(err))
		return err
	}

	logger.Info(c, "SHORTEN_LINK_SUCCESS", zap.String("link_info", utils.InterfaceToJsonString(link)))

	return c.JSON(link)
}
