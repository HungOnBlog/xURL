package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
	"hungon.space/xurl/common/utils"
	"hungon.space/xurl/common/validators"
)

type UserService struct {
	UserRepo *UserRepo
}

func (u *UserService) New() *UserService {
	return &UserService{
		UserRepo: &UserRepo{},
	}
}

func (u *UserService) generateUserID() string {
	return "user_" + utils.GenShortUUID()
}

func (u *UserService) genApikey() string {
	return utils.GenRandomString(32)
}

func (u *UserService) CreateUser(c *fiber.Ctx, userType string, email string, name string) error {
	if userType != "a" {
		isEmailValid := validators.IsEmailValid(email)
		if !isEmailValid {
			logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", xerror.EmailInvalid().Error()))
			return xerror.EmailInvalid()
		}
	}

	user := &User{
		UserId: u.generateUserID(),
		Type:   userType,
		ApiKey: u.genApikey(),
		Email:  email,
		Name:   name,
	}

	if err := u.UserRepo.CreateOne(user); err != nil {
		logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", err.Error()))
		return err
	}

	logger.Info(c, "CREATE_USER_SUCCESS", zap.String("data", utils.InterfaceToJsonString(user)))
	return c.JSON(user)
}
