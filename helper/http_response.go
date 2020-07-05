package helper

import (
	"net/http"
)

// HTTPResponse format
type HTTPResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// NewHTTPResponse for create common response
func NewHTTPResponse(code int, message string, data interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)

	commonResponse.Data = data

	if code < http.StatusBadRequest {
		commonResponse.Success = true
	}

	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}