package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"hungon.space/xurl/common/utils"
)

func ApplyRequestIdMiddleware(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Header: "requestId",
		Generator: func() string {
			return utils.GenShortUUID()
		},
	}))

	fmt.Println("RequestId middleware applied")
}
