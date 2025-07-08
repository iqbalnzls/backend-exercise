package rest

import "github.com/iqbalnzls/backend-exercise/internal/dto"

func toListingResponse(listing interface{}) dto.ListingResponse {
	return dto.ListingResponse{
		Result:  true,
		Listing: listing,
	}
}

func toUserResponse(user interface{}) dto.UserResponse {
	return dto.UserResponse{
		Result: true,
		User:   user,
	}
}
