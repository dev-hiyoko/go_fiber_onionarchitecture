package shared

import "net/http"

func GetErrorMessage(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "Bad request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not found"
	case http.StatusRequestTimeout:
		return "Request timeout"
	case http.StatusInternalServerError:
		return "Internal server error"
	default:
		return "Error"
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ValidateErrorResponse struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
