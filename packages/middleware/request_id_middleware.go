package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"hungon.space/xurl/common/utils"
)

func ApplyRequestIdMiddleware(a *fiber.App) {
	a.Use(requestid.New(requestid.Config{
		Header: "requestId",
		Generator: func() string {
			return utils.GenShortUUID()
		},
		ContextKey: "requestId",
	}))
	fmt.Println("Request ID middleware applied")
}
