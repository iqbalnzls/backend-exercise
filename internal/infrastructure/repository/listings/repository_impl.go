package listings

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
		panic("db is nil")
	}

	return &repo{
		db: db,
	}
}

func (r *repo) Save(logger *zap.Logger, listing *domain.Listings) (err error) {
	if err = r.db.Save(&listing).Error; err != nil {
		logger.Error("failed to save listing", zap.Error(err))
	}

	return
}

func (r *repo) FindAll(logger *zap.Logger, pageNum, pageSize int, args map[string]interface{}) (listings []domain.Listings, err error) {
	offset := (pageNum - 1) * pageSize

	if err = r.db.Preload("Users").Limit(pageSize).Offset(offset).Where(args).Find(&listings).Error; err != nil {
		logger.Error("failed to find listings", zap.Error(err))
		return
	}

	if len(listings) == 0 {
		err = errors.New("data not found")
		logger.Error("no listings found", zap.Error(err))
	}

	return
}
