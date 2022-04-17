package apiResponse

type TokensApiResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokensApiResponse(accessToken string, refreshToken string) *TokensApiResponse {
	return &TokensApiResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
