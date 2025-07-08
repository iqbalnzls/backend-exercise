package users

import (
	"errors"

	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/domain"
	"github.com/iqbalnzls/backend-exercise/pkg/database"
)

type repo struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	if db == nil {
		panic("database cannot be nil")
	}

	return &repo{
		db: db,
	}
}

func (r repo) Save(logger *zap.Logger, user *domain.Users) (err error) {
	if err = r.db.Save(&user).Error; err != nil {
		logger.Error("failed to save user", zap.Error(err))
	}
	return
}

func (r repo) FindBy(logger *zap.Logger, args map[string]interface{}) (user domain.Users, err error) {
	if err = r.db.Where(args).First(&user).Error; err != nil {
		logger.Error("failed to find user", zap.Error(err))
		return
	}

	return
}

func (r repo) FindAll(logger *zap.Logger, pageNum, pageSize int) (users []domain.Users, err error) {
	offset := (pageNum - 1) * pageSize

	if err = r.db.Limit(pageSize).Offset(offset).Order("created_at DESC").Find(&users).Error; err != nil {
		logger.Error("failed to find users", zap.Error(err))
		return
	}

	if len(users) == 0 {
		err = errors.New("data not found")
		logger.Error("no users found", zap.Error(err))
	}

	return
}
