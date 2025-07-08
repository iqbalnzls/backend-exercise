package rest

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/app/usecase/listings"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/pkg/validator"
)

type listingHandler struct {
	svc listings.Service
	v   *validator.Validator
}

func NewListingHandler(svc listings.Service, validator *validator.Validator) ListingHandler {
	if svc == nil {
		panic("service cannot be nil")

	}

	if validator == nil {
		panic("validator cannot be nil")

	}

	return &listingHandler{
		svc: svc,
		v:   validator,
	}
}

func (h *listingHandler) GetAll(c *fiber.Ctx) (err error) {
	req := &dto.GetAllListingsRequest{}
	logger := c.UserContext().Value("logger").(*zap.Logger)

	if err = c.QueryParser(req); err != nil {
		logger.Error("Failed to parse query parameters", zap.Error(err))
		return
	}

	if req.PageNum < 1 {
		req.PageNum = 1
	}

	if req.PageSize < 1 {
		req.PageSize = 10
	}

	getAllResp, err := h.svc.GetAll(logger, req)
	if err != nil {
		return
	}

	resp := toListingResponse(getAllResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *listingHandler) CreateListings(c *fiber.Ctx) (err error) {
	req := dto.CreateListingsRequest{}
	logger := c.UserContext().Value("logger").(*zap.Logger)

	if c.Get("Content-Type") != "application/x-www-form-urlencoded" {
		err = fiber.ErrUnsupportedMediaType
		logger.Error("Invalid content type, expected application/x-www-form-urlencoded", zap.Error(err))
		return
	}

	if err = c.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body", zap.Error(err))
		return
	}

	if err = h.v.Validate(req); err != nil {
		logger.Error("Validation failed", zap.Error(err))
		return
	}

	createResp, err := h.svc.Create(logger, &req)
	if err != nil {
		return
	}

	resp := toListingResponse(createResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)

}
