package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/di"
)

func StartHttpServer(container *di.Container) {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler(),
		AppName:      container.Config.App.Name,
	})

	app.Use(SetupMiddleware(container.Logger))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	SetupRouter(app, SetupHandler(container))

	app.Listen(container.Config.AppAddress())
}

func errorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		logger := ctx.UserContext().Value("logger").(*zap.Logger)

		logger.Info("Finishing request...")

		return ctx.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"result":  false,
				"message": err.Error(),
			})
	}
}
