package users

import (
	"bookstoreUsersApi/datasources/mysql/bookstores_users_db"
	"bookstoreUsersApi/utils/date_utils"
	"bookstoreUsersApi/utils/errors"
	"bookstoreUsersApi/utils/mysql_utils"
	"fmt"
)

const (
	QUERYINSERTUSER   = "INSERT INTO users(first_name, last_name, email, password, status, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	QUERYGETUSER      = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE id = ?"
	QUERYUPDATEUSER   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	QUERYDELETEUSER   = "DELETE FROM users WHERE id=?"
	QUERYFINDBYSTATUS = "SELECT id, first_name, last_name, email, password, status, created_at FROM users WHERE status=?;"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt); getErr != nil {
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

	user.CreatedAt = date_utils.GetNowDBFormat()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt)
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYFINDBYSTATUS)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = stmt.Close()
	}()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer func() {
		err = rows.Close()
	}()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt); err != nil {
			return nil, mysql_utils.ParsError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
