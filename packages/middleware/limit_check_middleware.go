package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"hungon.space/xurl/app/links"
	"hungon.space/xurl/app/users"
	"hungon.space/xurl/common/cache"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
	"hungon.space/xurl/common/utils"
)

func validatorChain(validators []func(*fiber.Ctx, *links.Link, *users.User) error, c *fiber.Ctx, l *links.Link, u *users.User) error {
	for _, validator := range validators {
		if err := validator(c, l, u); err != nil {
			return err
		}
	}
	return nil
}

func shortenTypeAValidator(c *fiber.Ctx, l *links.Link, u *users.User) error {
	if l.Type == "a" {
		return nil
	}

	return nil
}

func shortenTypePValidator(c *fiber.Ctx, l *links.Link, u *users.User) error {
	if l.Type == "p" {
		if u.LimitPassword <= 0 {
			return xerror.PasswordLimitReached()
		}

		if len(l.Password) < 9 {
			return xerror.PasswordInvalid()
		}

		return nil
	}

	return nil
}

func shortenTypeTValidator(c *fiber.Ctx, l *links.Link, u *users.User) error {
	if l.Type == "t" {
		if u.LimitTracking <= 0 {
			return xerror.TrackingLimitReached()
		}

		return nil
	}

	return nil
}

func shortenTypeTpValidator(c *fiber.Ctx, l *links.Link, u *users.User) error {
	if l.Type == "tp" {
		if u.LimitPassword <= 0 {
			return xerror.PasswordLimitReached()
		}

		if len(l.Password) < 9 {
			return xerror.PasswordInvalid()
		}

		if u.LimitTracking <= 0 {
			return xerror.TrackingLimitReached()
		}

		return nil
	}

	return nil
}

func updateLimit(c *fiber.Ctx, l *links.Link, u *users.User) {
	apikey := c.Get("apikey")
	cacheRepo := cache.RedisRepo{}

	if l.Type == "p" {
		u.LimitPassword = u.LimitPassword - 1
	}

	if l.Type == "t" {
		u.LimitTracking = u.LimitTracking - 1
	}

	if l.Type == "tp" {
		u.LimitTracking = u.LimitTracking - 1
		u.LimitPassword = u.LimitPassword - 1
	}

	userString, _ := json.Marshal(u)
	logger.Info(c, "UPDATE_LIMIT", zap.String("limit", string(userString)))
	cacheRepo.Set(apikey, userString)
}

func ApplyLimitCheckMiddleware(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		// Skip for /users paths
		if utils.StringInclude(c.Path(), "/users") {
			return c.Next()
		}

		logger.Info(c, "CHECK_LIMIT")
		userString := c.Get("user")
		if userString == "" {
			return c.JSON(xerror.UnauthorizedUserNotFound())
		}

		user := &users.User{}
		json.Unmarshal([]byte(userString), user)

		var body links.Link
		if err := c.BodyParser(&body); err != nil {
			return c.JSON(xerror.LinkBodyInvalid())
		}

		validators := []func(*fiber.Ctx, *links.Link, *users.User) error{
			shortenTypeAValidator,
			shortenTypePValidator,
			shortenTypeTValidator,
			shortenTypeTpValidator,
		}

		if err := validatorChain(validators, c, &body, user); err != nil {
			logger.Warn(c, "LIMIT_CHECK_REACHED", zap.String("error", err.Error()))
			return c.JSON(err)
		}

		logger.Info(c, "LIMIT_CHECK_PASSED", zap.String("limit", userString))
		updateLimit(c, &body, user)
		return c.Next()
	})
	fmt.Println("Limit check middleware applied")
}
