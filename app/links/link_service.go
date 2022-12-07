package links

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
	"hungon.space/xurl/common/utils"
)

type LinkServiceInterface interface {
	genLinkId() string
	ShortenLink(c *fiber.Ctx, originalLink string, apiKey string, password string, linkType string) error
	GetLinkTypeA(c *fiber.Ctx, linkId string) error
	GetLinkTypeP(c *fiber.Ctx, linkId string, password string) error
}

type LinkService struct {
	linkRepo LinkRepo
}

func (l *LinkService) New() *LinkService {
	return &LinkService{
		linkRepo: LinkRepo{},
	}
}

func (l *LinkService) genLinkId() string {
	id, _ := l.linkRepo.LastId()
	return utils.HashId(id + 1)
}

func (l *LinkService) ShortenLink(c *fiber.Ctx, originalLink string, apiKey string, password string, linkType string) error {
	linkId := l.genLinkId()
	fmt.Println(linkId)
	link := Link{
		LinkID:       linkId,
		OriginalLink: originalLink,
		ShortLink:    os.Getenv("BASE_URL") + "/" + linkType + "/" + linkId,
		ApiKey:       apiKey,
		Password:     password,
		Type:         linkType,
	}

	err := l.linkRepo.CreateOne(&link)

	if err != nil {
		logger.Warn(c, "SHORTEN_LINK_ERROR", zap.String("error", err.Error()))
		return xerror.InternalServerError()
	}

	logger.Info(c, "SHORTEN_LINK_SUCCESS", zap.String("data", utils.InterfaceToJsonString(link)))
	return c.JSON(link)
}

func (l *LinkService) GetLinkTypeA(c *fiber.Ctx, linkId string) error {
	var link Link
	err := l.linkRepo.FindBySelfID(linkId, &link)

	if err != nil {
		logger.Warn(c, "GET_LINK_ERROR", zap.String("error", err.Error()))
		return c.JSON(xerror.LinkNotFound())
	}

	return c.JSON(link)
}

func (l *LinkService) GetLinkTypeP(c *fiber.Ctx, linkId string, password string) error {
	if password == "" {
		return c.JSON(xerror.LinkPasswordRequired())
	}

	var link Link
	err := l.linkRepo.FindBySelfID(linkId, &link)

	if err != nil {
		logger.Warn(c, "GET_LINK_ERROR", zap.String("error", err.Error()))
		return c.JSON(xerror.LinkNotFound())
	}

	if link.Password != password {
		logger.Warn(c, "GET_LINK_ERROR", zap.String("error", xerror.LinkPasswordIncorrect().Error()))
		return c.JSON(xerror.LinkPasswordIncorrect())
	}

	return c.JSON(link)
}
