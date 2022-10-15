package response

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BuildSuccessResponse(code int, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  code,
		Message: message,
		Data:    data,
	}
}
