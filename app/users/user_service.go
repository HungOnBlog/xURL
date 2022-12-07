package users

import (
	"os"

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

func (u *UserService) freeUserTemplate(email string, password string, name string) *User {
	return &User{
		UserId:          u.generateUserID(),
		Type:            "f",
		ApiKey:          u.genApikey(),
		Email:           email,
		Name:            name,
		Password:        password,
		LimitPassword:   utils.StringToInt(os.Getenv("FREE_PASSWORD")),
		LimitTracking:   utils.StringToInt(os.Getenv("FREE_TRACKING")),
		LimitCustomSlug: utils.StringToInt(os.Getenv("FREE_SLUG")),
	}
}

func (u *UserService) proUserTemplate(email string, password string, name string) *User {
	return &User{
		UserId:          u.generateUserID(),
		Type:            "p",
		ApiKey:          u.genApikey(),
		Email:           email,
		Password:        password,
		Name:            name,
		LimitPassword:   utils.StringToInt(os.Getenv("PRO_PASSWORD")),
		LimitTracking:   utils.StringToInt(os.Getenv("PRO_TRACKING")),
		LimitCustomSlug: utils.StringToInt(os.Getenv("PRO_SLUG")),
	}
}

func (u *UserService) anonymousUserTemplate() *User {
	return &User{
		UserId:          u.generateUserID(),
		Type:            "a",
		ApiKey:          u.genApikey(),
		LimitPassword:   0,
		LimitTracking:   0,
		LimitCustomSlug: 0,
	}
}

func (u *UserService) CreateUser(c *fiber.Ctx, userType string, email string, name string, password string) error {
	if userType != "a" {
		isEmailValid := validators.IsEmailValid(email)
		if !isEmailValid {
			logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", xerror.EmailInvalid().Error()))
			return c.JSON(xerror.EmailInvalid())
		}

		if password == "" || len(password) < 9 {
			logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", xerror.PasswordInvalid().Error()))
			return c.JSON(xerror.PasswordInvalid())
		}
	}

	var user *User

	if userType == "f" {
		user = u.freeUserTemplate(email, password, name)
	} else if userType == "p" {
		user = u.proUserTemplate(email, password, name)
	} else {
		user = u.anonymousUserTemplate()
	}

	if err := u.UserRepo.CreateOne(user); err != nil {
		logger.Warn(c, "CREATE_USER_ERROR", zap.String("error", err.Error()))
		return c.JSON(err)
	}

	logger.Info(c, "CREATE_USER_SUCCESS", zap.String("data", utils.InterfaceToJsonString(user)))
	return c.JSON(user)
}

func (u *UserService) GetUserByApikey(c *fiber.Ctx, apikey string) (*User, error) {
	var user User
	err := u.UserRepo.FindByApiKey(apikey, &user)
	if err != nil {
		logger.Warn(c, "GET_USER_ERROR", zap.String("error", err.Error()))
		return nil, err
	}
	return &user, nil
}
