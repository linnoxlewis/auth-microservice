package apiResponse

type TokenApiResponse struct {
	Token  string `json:"token"`

}

func NewTokenApiResponse(token string) *TokenApiResponse{
	return &TokenApiResponse{token}
}
