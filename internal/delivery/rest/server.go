package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/iqbalnzls/backend-exercise/internal/di"
)

func StartHttpServer(app *fiber.App, container *di.Container) {
	app.Use(SetupMiddleware(container.Logger))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	SetupRouter(app, SetupHandler(container))

	app.Listen(container.Config.AppAddress())
}
