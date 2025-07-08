package users

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/internal/infrastructure/repository/users"
)

type service struct {
	userRepo users.Repository
}

func NewService(userRepo users.Repository) Service {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return &service{
		userRepo: userRepo,
	}

}

func (s *service) Create(logger *zap.Logger, req *dto.CreateUsersRequest) (resp dto.CreateUsersResponse, err error) {
	user := toUsers(req)

	if err = s.userRepo.Save(logger, user); err != nil {
		return
	}

	resp = toCreateUsersResponse(user)

	return
}

func (s *service) GetById(logger *zap.Logger, req *dto.GetByIdRequest) (resp dto.GetByIdResponse, err error) {
	args := map[string]interface{}{
		"id": req.Id,
	}

	user, err := s.userRepo.FindBy(logger, args)
	if err != nil {
		return
	}

	resp = toGetByIdResponse(&user)

	return
}

func (s *service) GetAll(logger *zap.Logger, req *dto.GetAllUsersRequest) (resp []dto.GetAllUsersResponse, err error) {
	usersData, err := s.userRepo.FindAll(logger, req.PageNum, req.PageSize)
	if err != nil {
		return
	}

	resp = toGetAllUsersResponse(usersData)

	return
}
