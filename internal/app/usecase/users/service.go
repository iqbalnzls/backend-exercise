package users

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
)

type Service interface {
	Create(logger *zap.Logger, req *dto.CreateUsersRequest) (resp dto.CreateUsersResponse, err error)
	GetById(logger *zap.Logger, req *dto.GetByIdRequest) (resp dto.GetByIdResponse, err error)
	GetAll(logger *zap.Logger, req *dto.GetAllUsersRequest) (resp []dto.GetAllUsersResponse, err error)
}
