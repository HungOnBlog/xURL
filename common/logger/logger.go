package logger

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger

func init() {
	lum := &lumberjack.Logger{
		Filename:   os.Getenv("LOG_PATH") + "/app.log",
		MaxSize:    500, // megabytes
		MaxAge:     7,   // days
		MaxBackups: 3,
	}
	w := zap.CombineWriteSyncers(os.Stdout, zapcore.AddSync(lum))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)

	zapLogger = zap.New(core)
}

func headerFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String("host", c.Hostname()),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.String("ip", c.IP()),
		zap.String("userAgent", c.Get("User-Agent")),
		zap.String("requestId", c.Context().Value("requestId").(string)),
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
