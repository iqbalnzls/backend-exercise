package listings

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/domain"
)

type Repository interface {
	Save(logger *zap.Logger, listing *domain.Listings) error
	FindAll(logger *zap.Logger, pageNum, pageSize int, args map[string]interface{}) ([]domain.Listings, error)
}
