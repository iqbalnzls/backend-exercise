package listings

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
)

type Service interface {
	Create(logger *zap.Logger, req *dto.CreateListingsRequest) (dto.CreateListingsResponse, error)
	GetAll(logger *zap.Logger, req *dto.GetAllListingsRequest) ([]dto.GetAllListingsResponse, error)
}
