package middleware

import "github.com/gofiber/fiber/v2"

func AppMiddlewares(a *fiber.App) {
	ApplyRequestIdMiddleware(a)
}
