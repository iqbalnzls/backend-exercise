package public_api

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
)

type Service interface {
	GetAllListingsPublicApi(logger *zap.Logger, req *dto.GetAllListingsPublicApiRequest) (resp []dto.GetAllListingsPublicApiResponse, err error)
	CreateUserPublicApi(logger *zap.Logger, req *dto.CreateUserPublicApiRequest) (resp dto.CreateUserPublicApiResponse, err error)
	CreateListingsPublicApi(logger *zap.Logger, req *dto.CreateListingsPublicApiRequest) (resp dto.CreateListingsPublicApiResponse, err error)
}
