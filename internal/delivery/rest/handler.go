package rest

import (
	"github.com/gofiber/fiber/v2"

	"github.com/iqbalnzls/backend-exercise/internal/di"
)

type (
	ListingHandler interface {
		Save(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
	}

	UserHandler interface {
		Save(c *fiber.Ctx) error
		GetById(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
	}
)

type Handler struct {
	listingHandler ListingHandler
	userHandler    UserHandler
}

func SetupHandler(container *di.Container) *Handler {
	listing := NewListingHandler(container.ListingService, container.Validator)
	user := NewUserHandler(container.UserService, container.Validator)

	return &Handler{
		listingHandler: listing,
		userHandler:    user,
	}

}
