package apiResponse

type TokensApiResponse struct {
	AccessToken  string
	RefreshToken string
}

func NewTokensApiResponse(accessToken string,refreshToken string) *TokensApiResponse{
	return &TokensApiResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}
