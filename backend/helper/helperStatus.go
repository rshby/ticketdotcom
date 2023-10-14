package helper

import "net/http"

func GenerateStatusFromCode(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "ok"
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusNotFound:
		return "not found"
	case http.StatusUnauthorized:
		return "unauthorized"
	default:
		return "internal server error"
	}
}
