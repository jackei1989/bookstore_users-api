package users

import (
	"bookstoreUsersApi/utils/errors"
	"strings"
)

var (
	StatusActive = "active"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	// if user.FirstName == "" {
	// 	return errors.New("first name is required")
	// }
	// if user.LastName == "" {
	// 	return errors.New("last name is required")
	// }
	// if user.Email == "" {
	// 	return errors.New("email is required")
	// }

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
