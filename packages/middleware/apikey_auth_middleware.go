package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"hungon.space/xurl/app/users"
	xerror "hungon.space/xurl/common/error"
	"hungon.space/xurl/common/logger"
)

func ApplyApikeyAuthMiddleware(a *fiber.App) {
	userService := &users.UserService{}
	a.Use(func(c *fiber.Ctx) error {
		// Skip for /users paths
		if c.Path() == "/users" || c.Path() == "/users/" {
			return c.Next()
		}

		apiKey := c.Get("apiKey")

		if apiKey == "" {
			logger.Warn(c, "AUTHENTICATION", zap.String("error", xerror.ApikeyInvalid().Error()))
			return c.JSON(xerror.ApikeyInvalid())
		}

		logger.Info(c, "GET_USER")
		user, err := userService.GetUserByApikey(c, apiKey)
		if err != nil {
			return c.JSON(xerror.UnauthorizedUserNotFound())
		}

		userString, _ := json.Marshal(user)
		c.Request().Header.Set("user", string(userString))

		return c.Next()
	})
	fmt.Println("API Key middleware applied")
}
