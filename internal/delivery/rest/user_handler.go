package rest

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/app/usecase/users"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/pkg/validator"
)

type userHandler struct {
	svc users.Service
	v   *validator.Validator
}

func NewUserHandler(svc users.Service, validator *validator.Validator) UserHandler {
	if svc == nil {
		panic("service cannot be nil")

	}

	if validator == nil {
		panic("validator cannot be nil")

	}

	return &userHandler{
		svc: svc,
		v:   validator,
	}
}

func (h *userHandler) CreateListings(c *fiber.Ctx) (err error) {
	req := dto.CreateUsersRequest{}
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

	resp := toUserResponse(createResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *userHandler) GetById(c *fiber.Ctx) (err error) {
	logger := c.UserContext().Value("logger").(*zap.Logger)

	userId, err := c.ParamsInt("id")
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		return
	}

	req := &dto.GetByIdRequest{
		Id: int64(userId),
	}

	getByIdResp, err := h.svc.GetById(logger, req)
	if err != nil {
		return
	}

	resp := toUserResponse(getByIdResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *userHandler) GetAll(c *fiber.Ctx) (err error) {
	logger := c.UserContext().Value("logger").(*zap.Logger)
	req := &dto.GetAllUsersRequest{}

	if err = c.QueryParser(req); err != nil {
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

	resp := toUserResponse(getAllResp)

	logger.Info("Finishing request...",
		zap.Any("resp", resp))

	return c.Status(fiber.StatusOK).JSON(resp)
}
