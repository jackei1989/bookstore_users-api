package users

import (
	"bookstoreUsersApi/datasources/mysql/bookstores_users_db"
	"bookstoreUsersApi/utils/date_utils"
	"bookstoreUsersApi/utils/errors"
	"fmt"
	"strings"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	QUERYINSERTUSER  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	INDEXUNIQUEEMAIL = "email"
)

func (user *User) Get() *errors.RestErr {

	if err := bookstores_users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYINSERTUSER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
	}()

	user.CreatedAt = date_utils.GetNowString()

	fmt.Println(user.CreatedAt)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), INDEXUNIQUEEMAIL) {
			return errors.NewInternalServerError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when try to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when try to save user: %s", err.Error()))
	}
	user.Id = userId
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.CreatedAt = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
