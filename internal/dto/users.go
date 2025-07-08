package dto

type CreateUsersRequest struct {
	Name string `form:"name" validate:"required"`
}

type CreateUsersResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserResponse struct {
	Result bool        `json:"result"`
	User   interface{} `json:"user"`
}

type GetByIdRequest struct {
	Id int64
}

type GetByIdResponse struct {
	CreateUsersResponse
}

type GetAllUsersRequest struct {
	PageNum  int `query:"page_num"`
	PageSize int `query:"page_size"`
}

type GetAllUsersResponse struct {
	CreateUsersResponse
}
