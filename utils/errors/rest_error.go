package errors

import (
	"net/http"
)

type RestErr struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Message: message,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
		Message: message,
	}
}
