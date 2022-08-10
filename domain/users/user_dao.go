package users

import (
	"bookstoreUsersApi/datasources/mysql/bookstores_users_db"
	"bookstoreUsersApi/utils/date_utils"
	"bookstoreUsersApi/utils/errors"
	"fmt"
	"strings"
)

const (
	QUERYINSERTUSER  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	QUERYGETUSER     = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id = ?"
	INDEXUNIQUEEMAIL = "email"
	ERRORNOROWS      = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYGETUSER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
	}()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), ERRORNOROWS) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found!", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when try to get user %d: %s", user.Id, err.Error()))
	}
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
	return nil
}
