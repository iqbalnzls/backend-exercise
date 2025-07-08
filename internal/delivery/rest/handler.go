package rest

import (
	"github.com/gofiber/fiber/v2"

	"github.com/iqbalnzls/backend-exercise/internal/di"
)

type (
	ListingHandler interface {
		CreateListings(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
	}

	UserHandler interface {
		CreateListings(c *fiber.Ctx) error
		GetById(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
	}

	PublicApiHandler interface {
		GetAllListingsPublicApi(c *fiber.Ctx) error
		CreateListingsPublicApi(c *fiber.Ctx) error
		CreateUserPublicApi(c *fiber.Ctx) error
	}
)

type Handler struct {
	listingHandler   ListingHandler
	userHandler      UserHandler
	publicApiHandler PublicApiHandler
}

func SetupHandler(container *di.Container) *Handler {
	listing := NewListingHandler(container.ListingService, container.Validator)
	user := NewUserHandler(container.UserService, container.Validator)
	publicApi := NewPublicApiHandler(container.PublicApiService, container.Validator)

	return &Handler{
		listingHandler:   listing,
		userHandler:      user,
		publicApiHandler: publicApi,
	}

}
