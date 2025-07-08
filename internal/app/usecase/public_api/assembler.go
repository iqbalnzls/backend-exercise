package public_api

import (
	"github.com/iqbalnzls/backend-exercise/internal/domain"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/pkg/constant"
)

func toGetAllListingsPublicResponses(listings []domain.Listings) []dto.GetAllListingsPublicApiResponse {
	res := make([]dto.GetAllListingsPublicApiResponse, 0)

	for _, v := range listings {
		res = append(res, toGetAllListingsPublicResponse(v))
	}

	return res
}

func toGetAllListingsPublicResponse(listing domain.Listings) dto.GetAllListingsPublicApiResponse {
	return dto.GetAllListingsPublicApiResponse{
		Id:          listing.Id,
		ListingType: constant.ListingType(listing.ListingType),
		Price:       listing.Price,
		CreatedAt:   listing.CreatedAt,
		UpdatedAt:   listing.UpdatedAt,
		User:        toCreateUsersResponse(listing.Users),
	}
}

func toCreateUsersResponse(user domain.Users) dto.CreateUsersResponse {
	return dto.CreateUsersResponse{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toListings(req *dto.CreateListingsPublicApiRequest) *domain.Listings {
	return &domain.Listings{
		UserId:      req.UserId,
		Price:       req.Price,
		ListingType: string(req.ListingType),
	}
}

func toCreateListingsPublicApiResponse(listing *domain.Listings) dto.CreateListingsPublicApiResponse {
	return dto.CreateListingsPublicApiResponse{
		CreateListingsResponse: dto.CreateListingsResponse{
			Id:          listing.Id,
			UserId:      listing.UserId,
			ListingType: constant.ListingType(listing.ListingType),
			Price:       listing.Price,
			CreatedAt:   listing.CreatedAt,
			UpdatedAt:   listing.UpdatedAt,
		},
	}
}

func toUsers(req *dto.CreateUserPublicApiRequest) *domain.Users {
	return &domain.Users{
		Name: req.Name,
	}
}

func toCreateUserPublicApiResponse(user *domain.Users) dto.CreateUserPublicApiResponse {
	return dto.CreateUserPublicApiResponse{
		CreateUsersResponse: dto.CreateUsersResponse{
			Id:        user.Id,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
