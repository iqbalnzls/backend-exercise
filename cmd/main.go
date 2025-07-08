package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/delivery/rest"
	"github.com/iqbalnzls/backend-exercise/internal/di"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			logger := ctx.UserContext().Value("logger").(*zap.Logger)

			logger.Info("Finishing request...")

			return ctx.Status(fiber.StatusOK).JSON(
				fiber.Map{
					"result":  false,
					"message": err.Error(),
				})
		},
	})

	container := di.SetupContainer()

	rest.StartHttpServer(app, container)
}
