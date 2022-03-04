package apiResponse

type SuccessApiResponse struct {
	Success bool `json:"success"`
	Data string `json:"data"`
}

func NewSuccessApiResponse(success bool,data string) *SuccessApiResponse {
	return &SuccessApiResponse{success,data}
}