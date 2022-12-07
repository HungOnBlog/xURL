package users

import (
	"time"

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
	return utils.GenApikey("")
}

func (u *UserService) CreateUser(c *fiber.Ctx, userType string, email string, name string, password string) error {
	if userType != "a" {
		isEmailValid := validators.IsEmailValid(email)
		if !isEmailValid {
			logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", xerror.EmailInvalid().Error()))
			return xerror.EmailInvalid()
		}

		if password == "" || len(password) < 9 {
			logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", xerror.PasswordInvalid().Error()))
			return xerror.PasswordInvalid()
		}
	}

	var expiredDate time.Time
	// !NOTE: This setting can be affected by Y2K38 problem.
	if userType == "a" {
		expiredDate = time.Now().AddDate(0, 0, 1) // Anonymous user will be expired after 1 day
	} else {
		expiredDate = time.Now().AddDate(100, 0, 0) // Normal user will be expired after 100 years.
	}

	user := &User{
		UserId:      u.generateUserID(),
		Type:        userType,
		ApiKey:      u.genApikey(),
		Email:       email,
		Name:        name,
		Password:    password,
		ExpiredDate: expiredDate,
	}

	if err := u.UserRepo.CreateOne(user); err != nil {
		logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", err.Error()))
		return err
	}

	logger.Info(c, "CREATE_USER_SUCCESS", zap.String("data", utils.InterfaceToJsonString(user)))
	return c.JSON(user)
}
