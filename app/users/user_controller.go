package users

import (
	"github.com/gofiber/fiber/v2"
	xerror "hungon.space/xurl/common/error"

	"hungon.space/xurl/common/logger"
)

func UserController(r fiber.Router) {
	userService := UserService{}

	// User routes
	// CREATE USER
	r.Post("/users", func(c *fiber.Ctx) error {
		logger.Info(c, "CREATE_USER")
		var user User
		err := c.BodyParser(&user)
		if err != nil {
			return xerror.UserBodyInvalid()
		}

		return userService.CreateUser(c, user.Type, user.Email, user.Name, user.Password)
	})
}
