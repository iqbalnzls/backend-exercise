package listings

import (
	"github.com/iqbalnzls/backend-exercise/internal/domain"
	"github.com/iqbalnzls/backend-exercise/internal/dto"
	"github.com/iqbalnzls/backend-exercise/pkg/constant"
)

func toListings(req *dto.CreateListingsRequest) *domain.Listings {
	return &domain.Listings{
		UserId:      req.UserId,
		Price:       req.Price,
		ListingType: string(req.ListingType),
	}
}

func toCreateListingsResponse(listings *domain.Listings) dto.CreateListingsResponse {
	return dto.CreateListingsResponse{
		Id:          listings.Id,
		UserId:      listings.UserId,
		ListingType: constant.ListingType(listings.ListingType),
		Price:       listings.Price,
		CreatedAt:   listings.CreatedAt,
		UpdatedAt:   listings.UpdatedAt,
	}
}

func toGetAllListingsResponses(listings []domain.Listings) []dto.GetAllListingsResponse {
	res := make([]dto.GetAllListingsResponse, 0)

	for _, v := range listings {
		res = append(res, toGetAllListingsResponse(v))
	}

	return res
}

func toGetAllListingsResponse(listing domain.Listings) dto.GetAllListingsResponse {
	return dto.GetAllListingsResponse{
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
