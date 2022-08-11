package services

import (
	"bookstoreUsersApi/domain/users"
	"bookstoreUsersApi/utils/crypto_utils"
	"bookstoreUsersApi/utils/date_utils"
	"bookstoreUsersApi/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	Get(int64) (*users.User, *errors.RestErr)
	Create(users.User) (*users.User, *errors.RestErr)
	Update(bool, users.User) (*users.User, *errors.RestErr)
	Delete(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *userService) Get(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) Create(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.CreatedAt = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Update(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.Get(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) Delete(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
