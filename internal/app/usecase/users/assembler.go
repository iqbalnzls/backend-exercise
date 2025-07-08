package users

import (
	"github.com/iqbalnzls/backend-exercise/internal/domain"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
)

func toUsers(req *dto.CreateUsersRequest) *domain.Users {
	return &domain.Users{
		Name: req.Name,
	}
}

func toCreateUsersResponse(user *domain.Users) dto.CreateUsersResponse {
	return dto.CreateUsersResponse{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toGetByIdResponse(user *domain.Users) dto.GetByIdResponse {
	return dto.GetByIdResponse{
		CreateUsersResponse: dto.CreateUsersResponse{
			Id:        user.Id,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}

func toGetAllUsersResponse(users []domain.Users) []dto.GetAllUsersResponse {
	res := make([]dto.GetAllUsersResponse, 0)

	for _, v := range users {
		res = append(res, toGetAllUserResponse(v))
	}

	return res
}

func toGetAllUserResponse(user domain.Users) dto.GetAllUsersResponse {
	return dto.GetAllUsersResponse{
		CreateUsersResponse: dto.CreateUsersResponse{
			Id:        user.Id,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
