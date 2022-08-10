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
	QUERYUPDATEUSER = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	QUERYDELETEUSER = "DELETE FROM users WHERE id=?"
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

func (user *User) Update() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYUPDATEUSER)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
	}()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParsError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYDELETEUSER)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
	}()
	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParsError(err)
	}
	return nil
}
