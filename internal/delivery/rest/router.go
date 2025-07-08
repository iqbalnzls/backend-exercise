package rest

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, handler *Handler) {
	//health check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("service is up and running...")
	})

	// listing routes
	listings := app.Group("/listings")
	{
		listings.Use(AllowlistIPMiddleware(), BasicAuthMiddleware())
		listings.Post("", handler.listingHandler.CreateListings)
		listings.Get("", handler.listingHandler.GetAll)
	}

	// user routes
	users := app.Group("/users")
	{
		users.Use(AllowlistIPMiddleware(), BasicAuthMiddleware())
		users.Post("", handler.userHandler.CreateListings)
		users.Get("/:id", handler.userHandler.GetById)
		users.Get("", handler.userHandler.GetAll)
	}

	//public api routes
	publicApi := app.Group("/public-api")
	{
		publicApi.Get("/listings", handler.publicApiHandler.GetAllListingsPublicApi)
		publicApi.Post("/listings", handler.publicApiHandler.CreateListingsPublicApi)
		publicApi.Post("/users", handler.publicApiHandler.CreateUserPublicApi)
	}

}
