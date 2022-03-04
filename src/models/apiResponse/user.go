package apiResponse

type UserApiResponse struct {
	Id uint `json:"userId"`
	Email string `json:"email"`
}

func NewUserApiResponse(id uint,email string) *UserApiResponse {
	return &UserApiResponse{id,email}
}