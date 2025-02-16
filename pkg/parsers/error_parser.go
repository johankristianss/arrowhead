package parsers

import "encoding/json"

type ErrorResponse struct {
	ErrorMessage  string `json:"errorMessage"`
	ErrorCode     int    `json:"errorCode"`
	ExceptionType string `json:"exceptionType"`
}

func UnmarshalErrorResponse(body []byte) ErrorResponse {
	var errorResponse ErrorResponse
	json.Unmarshal(body, &errorResponse)
	return errorResponse
}
