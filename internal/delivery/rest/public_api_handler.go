package rest

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	publicApiService "github.com/iqbalnzls/backend-exercise/internal/app/usecase/public_api"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/pkg/validator"
)

type publicApiHandler struct {
	publicApiSvc publicApiService.Service
	v            *validator.Validator
}

func NewPublicApiHandler(publicApiSvc publicApiService.Service, v *validator.Validator) PublicApiHandler {
	if publicApiSvc == nil {
		panic("public api service cannot be nil")
	}

	if v == nil {
		panic("validator cannot be nil")
	}

	return &publicApiHandler{
		publicApiSvc: publicApiSvc,
		v:            v,
	}
}

func (h *publicApiHandler) GetAllListingsPublicApi(c *fiber.Ctx) (err error) {
	req := &dto.GetAllListingsPublicApiRequest{}
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

	getAllResp, err := h.publicApiSvc.GetAllListingsPublicApi(logger, req)
	if err != nil {
		return
	}

	resp := toListingResponse(getAllResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *publicApiHandler) CreateListingsPublicApi(c *fiber.Ctx) (err error) {
	req := &dto.CreateListingsPublicApiRequest{}
	logger := c.UserContext().Value("logger").(*zap.Logger)

	if c.Get("Content-Type") != "application/json" {
		err = fiber.ErrUnsupportedMediaType
		logger.Error("Invalid content type, expected application/json", zap.Error(err))
		return
	}

	if err = c.BodyParser(req); err != nil {
		logger.Error("Failed to parse request body", zap.Error(err))
		return
	}

	if err = h.v.Validate(req); err != nil {
		logger.Error("Validation failed", zap.Error(err))
		return

	}

	createListing, err := h.publicApiSvc.CreateListingsPublicApi(logger, req)
	if err != nil {
		return
	}

	resp := toListingResponse(createListing)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *publicApiHandler) CreateUserPublicApi(c *fiber.Ctx) (err error) {
	req := &dto.CreateUserPublicApiRequest{}
	logger := c.UserContext().Value("logger").(*zap.Logger)

	if c.Get("Content-Type") != "application/json" {
		err = fiber.ErrUnsupportedMediaType
		logger.Error("Invalid content type, expected application/json", zap.Error(err))
		return
	}

	if err = c.BodyParser(req); err != nil {
		logger.Error("Failed to parse request body", zap.Error(err))
		return
	}

	if err = h.v.Validate(req); err != nil {
		logger.Error("Validation failed", zap.Error(err))
		return
	}

	createUser, err := h.publicApiSvc.CreateUserPublicApi(logger, req)
	if err != nil {
		return
	}

	resp := toUserResponse(createUser)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}
