package dto

import "github.com/iqbalnzls/backend-exercise/pkg/constant"

type CreateListingsRequest struct {
	UserId      int64                `form:"user_id" validate:"required"`
	ListingType constant.ListingType `form:"listing_type" validate:"required,oneof=rent sale"` // Assuming these are the only valid types
	Price       int64                `form:"price" validate:"required"`
}

type CreateListingsResponse struct {
	Id          int64                `json:"id"`
	UserId      int64                `json:"user_id"`
	ListingType constant.ListingType `json:"listing_type"`
	Price       int64                `json:"price"`
	CreatedAt   int64                `json:"created_at"`
	UpdatedAt   int64                `json:"updated_at"`
}

type ListingResponse struct {
	Result  bool        `json:"result"`
	Listing interface{} `json:"listing"`
}

type GetAllListingsRequest struct {
	PageNum  int    `query:"page_num"`
	PageSize int    `query:"page_size"`
	UserId   string `query:"user_id"`
}

type GetAllListingsResponse struct {
	CreateListingsResponse
}
