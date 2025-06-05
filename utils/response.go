package utils

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
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

func FailedResponseWithData(message string, data []error) ErrorResponse {
	var errorStrings []string
	for _, err := range data {
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}
	return ErrorResponse{
		Status:  "failed",
		Message: message,
		Errors:  errorStrings,
	}
}