package users

import (
	"bookstoreUsersApi/datasources/mysql/bookstores_users_db"
	"bookstoreUsersApi/utils/date_utils"
	"bookstoreUsersApi/utils/errors"
	"bookstoreUsersApi/utils/mysql_utils"
	"fmt"
)

const (
	QUERYINSERTUSER = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	QUERYGETUSER    = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id = ?"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); getErr != nil {
		return mysql_utils.ParsError(getErr)
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if saveErr != nil {
		return mysql_utils.ParsError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParsError(saveErr)
	}
	user.Id = userId
	return nil
}
