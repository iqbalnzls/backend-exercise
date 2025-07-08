package listings

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
	listingsRepository "github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/listings"
	"github.com/iqbalnzls/backend-exercise/pkg/strutil"
)

type service struct {
	listingsRepo listingsRepository.Repository
}

func NewService(listingsRepo listingsRepository.Repository) Service {
	if listingsRepo == nil {
		panic("listingsRepo is nil")
	}

	return &service{
		listingsRepo: listingsRepo,
	}
}

func (s *service) Create(logger *zap.Logger, req *dto.CreateListingsRequest) (resp dto.CreateListingsResponse, err error) {
	listings := toListings(req)

	if err = s.listingsRepo.Save(logger, listings); err != nil {
		return
	}

	resp = toCreateListingsResponse(listings)

	return
}

func (s *service) GetAll(logger *zap.Logger, req *dto.GetAllListingsRequest) (resp []dto.GetAllListingsResponse, err error) {
	args := map[string]interface{}{}

	if !strutil.IsEmptyString(req.UserId) {
		args = map[string]interface{}{
			"user_id": req.UserId,
		}
	}

	listings, err := s.listingsRepo.FindAll(logger, req.PageNum, req.PageSize, args)
	if err != nil {
		return
	}

	resp = toGetAllListingsResponses(listings)

	return
}
