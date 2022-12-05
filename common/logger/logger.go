package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	zapLogger, _ = zap.NewProduction()
}

func headerFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String("host", c.Hostname()),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.String("ip", c.IP()),
		zap.String("userAgent", c.Get("User-Agent")),
		zap.String("requestId", c.GetReqHeaders()["requestId"]),
	}
}

func bodyFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String("body", string(c.Body())),
	}
}

func headerAndBodyFields(c *fiber.Ctx) []zap.Field {
	return append(headerFields(c), bodyFields(c)...)
}

func Info(c *fiber.Ctx, msg string, addition ...zap.Field) {
	zapLogger.Info(msg, append(
		headerAndBodyFields(c),
		addition...,
	)...)
}

func Error(c *fiber.Ctx, msg string, addition ...zap.Field) {
	zapLogger.Error(msg, append(
		headerAndBodyFields(c),
		addition...,
	)...)
}

func Fatal(c *fiber.Ctx, msg string, addition ...zap.Field) {
	zapLogger.Fatal(msg, append(
		headerAndBodyFields(c),
		addition...,
	)...)
}

func Warn(c *fiber.Ctx, msg string, addition ...zap.Field) {
	zapLogger.Warn(msg, append(
		headerAndBodyFields(c),
		addition...,
	)...)
}
