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