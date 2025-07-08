package dto

import "github.com/iqbalnzls/backend-exercise/pkg/constant"

type GetAllListingsPublicApiRequest struct {
	GetAllListingsRequest
}

type GetAllListingsPublicApiResponse struct {
	Id          int64                `json:"id"`
	ListingType constant.ListingType `json:"listing_type"`
	Price       int64                `json:"price"`
	CreatedAt   int64                `json:"created_at"`
	UpdatedAt   int64                `json:"updated_at"`
	User        CreateUsersResponse  `json:"user"`
}

type CreateUserPublicApiRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateUserPublicApiResponse struct {
	CreateUsersResponse
}

type CreateListingsPublicApiRequest struct {
	UserId      int64                `json:"user_id" validate:"required"`
	ListingType constant.ListingType `json:"listing_type" validate:"required,oneof=rent sale"`
	Price       int64                `json:"price" validate:"required"`
}

type CreateListingsPublicApiResponse struct {
	CreateListingsResponse
}
