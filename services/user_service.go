package services

import (
	"bookstoreUsersApi/domain/users"
	"bookstoreUsersApi/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, &errors.RestErr{
		Status:  http.StatusInternalServerError,
		Error:   "500",
		Message: "oops! something went wrong",
	}
}
