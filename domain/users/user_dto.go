package users

import (
	"bookstoreUsersApi/utils/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	// if user.FirstName == "" {
	// 	return errors.New("first name is required")
	// }
	// if user.LastName == "" {
	// 	return errors.New("last name is required")
	// }
	// if user.Email == "" {
	// 	return errors.New("email is required")
	// }
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
