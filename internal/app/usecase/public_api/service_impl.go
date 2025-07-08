package public_api

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
	listingsRepository "github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/listings"
	"github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/users"
	"github.com/iqbalnzls/backend-exercise/pkg/strutil"
)

type service struct {
	listingRepo listingsRepository.Repository
	userRepo    users.Repository
}

func NewService(listingRepo listingsRepository.Repository, userRepo users.Repository) Service {
	if listingRepo == nil {
		panic("listingRepo is nil")
	}
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return &service{
		listingRepo: listingRepo,
		userRepo:    userRepo,
	}
}

func (s *service) GetAllListingsPublicApi(logger *zap.Logger, req *dto.GetAllListingsPublicApiRequest) (resp []dto.GetAllListingsPublicApiResponse, err error) {
	args := make(map[string]interface{})
	if !strutil.IsEmptyString(req.UserId) {
		args = map[string]interface{}{
			"user_id": req.UserId,
		}
	}

	listings, err := s.listingRepo.FindAll(logger, req.PageNum, req.PageSize, args)
	if err != nil {
		return
	}

	resp = toGetAllListingsPublicResponses(listings)
	return
}

func (s *service) CreateListingsPublicApi(logger *zap.Logger, req *dto.CreateListingsPublicApiRequest) (resp dto.CreateListingsPublicApiResponse, err error) {
	listings := toListings(req)
	if err = s.listingRepo.Save(logger, listings); err != nil {
		return
	}

	resp = toCreateListingsPublicApiResponse(listings)
	return
}

func (s *service) CreateUserPublicApi(logger *zap.Logger, req *dto.CreateUserPublicApiRequest) (resp dto.CreateUserPublicApiResponse, err error) {
	user := toUsers(req)
	if err = s.userRepo.Save(logger, user); err != nil {
		return
	}

	resp = toCreateUserPublicApiResponse(user)
	return
}
