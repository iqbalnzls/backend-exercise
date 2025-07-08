package rest

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go.uber.org/zap"
)

// AllowlistIPMiddleware checks if the request comes from an allowed IP address.
func AllowlistIPMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		logger := ctx.UserContext().Value("logger").(*zap.Logger)

		ip := ctx.Get("Ip-Address")
		if ip != "127.192.1.1" {
			err = fiber.ErrForbidden
			logger.Error("Forbidden request: unauthorized IP address", zap.String("ip", ip), zap.Error(err))
			return
		}

		return ctx.Next()
	}
}

func BasicAuthMiddleware() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "password123",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			logger := c.UserContext().Value("logger").(*zap.Logger)

			logger.Error("Unauthorized access attempt", zap.Error(errors.New("unauthorized access")))

			return fiber.ErrForbidden
		},
	})
}

func SetupMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Info("Incoming request...",
			zap.String("uri", c.Request().URI().String()),
			zap.String("method", c.Method()),
			zap.Any("req", c.Body()))

		c.SetUserContext(context.WithValue(context.Background(), "logger", logger))

		return c.Next()
	}
}
