package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zapLogger = logger
}

func headerFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String("requestId", c.GetReqHeaders()["requestId"]),
		zap.String("userAgent", c.GetReqHeaders()["user-agent"]),
		zap.String("ip", c.IP()),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
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

func Info(c *fiber.Ctx, msg string, fields ...zap.Field) {
	zapLogger.Info(msg, append(headerAndBodyFields(c), fields...)...)
}

func Debug(c *fiber.Ctx, msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, append(headerAndBodyFields(c), fields...)...)
}

func Warn(c *fiber.Ctx, msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, append(headerAndBodyFields(c), fields...)...)
}

func Error(c *fiber.Ctx, msg string, fields ...zap.Field) {
	zapLogger.Error(msg, append(headerAndBodyFields(c), fields...)...)
}

func Fatal(c *fiber.Ctx, msg string, fields ...zap.Field) {
	zapLogger.Fatal(msg, append(headerAndBodyFields(c), fields...)...)
}
