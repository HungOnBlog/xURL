package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"hungon.space/xurl/app/users"
	"hungon.space/xurl/common/cache"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
	"hungon.space/xurl/common/utils"
)

func ApplyApikeyAuthMiddleware(a *fiber.App) {
	userService := &users.UserService{}
	a.Use(func(c *fiber.Ctx) error {
		cacheRepo := cache.RedisRepo{}
		// Skip for /users paths
		if utils.StringInclude(c.Path(), "/users") {
			return c.Next()
		}

		apiKey := c.Get("apiKey")

		if apiKey == "" {
			logger.Warn(c, "AUTHENTICATION", zap.String("error", xerror.ApikeyInvalid().Error()))
			return c.JSON(xerror.ApikeyInvalid())
		}

		userString, err := cacheRepo.Get(apiKey)

		if err == nil {
			logger.Info(c, "GET_USER", zap.String("user", string(userString)))
			c.Request().Header.Set("user", string(userString))
			return c.Next()
		} else {
			userInfo, err := userService.GetUserByApikey(c, apiKey)
			if err != nil {
				return c.JSON(xerror.UnauthorizedUserNotFound())
			}

			userString, _ := json.Marshal(userInfo)
			err = cacheRepo.Setex(apiKey, userString, time.Hour*24)
			if err != nil {
				logger.Warn(c, "CACHE_ERROR", zap.String("error", err.Error()))
			}
			c.Request().Header.Set("user", string(userString))

			return c.Next()
		}
	})
	fmt.Println("API Key middleware applied")
}
