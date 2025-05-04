package utils

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse creates a success response with the given data.
func SuccessResponse(data interface{}) Response {
	return Response{
		Status:  "success",
		Message: "success",
		Data:    data,
	}
}

// FailedResponse creates a failed response with the given message.
func FailedResponse(message string) Response {
	return Response{
		Status:  "failed",
		Message: message,
	}
}