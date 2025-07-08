package users

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/domain"
)

type Repository interface {
	Save(logger *zap.Logger, user *domain.Users) error
	FindBy(logger *zap.Logger, args map[string]interface{}) (domain.Users, error)
	FindAll(logger *zap.Logger, pageNum, pageSize int) ([]domain.Users, error)
}
