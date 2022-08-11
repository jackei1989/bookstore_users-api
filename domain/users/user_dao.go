package users

import (
	"bookstoreUsersApi/datasources/mysql/bookstores_users_db"
	"bookstoreUsersApi/logger"
	"bookstoreUsersApi/utils/errors"
	"fmt"
)

const (
	QUERYGETUSER      = "SELECT id, first_name, last_name, email,password, status, created_at FROM users WHERE id = ?"
	QUERYINSERTUSER   = "INSERT INTO users(first_name, last_name, email, password, status, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	QUERYUPDATEUSER   = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	QUERYDELETEUSER   = "DELETE FROM users WHERE id=?"
	QUERYFINDBYSTATUS = "SELECT id, first_name, last_name, email, password, status, created_at FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYGETUSER)
	if err != nil {
		logger.Error("error when trying to perpare get user statement", err)
		return errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = stmt.Close()
	}()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return errors.NewInternalServerError("Database Error!")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYINSERTUSER)
	if err != nil {
		logger.Error("error when trying to perpare save user statement", err)
		return errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = stmt.Close()
	}()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("Database Error!")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert user id after creating a new user", err)
		return errors.NewInternalServerError("Database Error!")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYUPDATEUSER)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = stmt.Close()
	}()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("Database Error!")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYDELETEUSER)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = stmt.Close()
	}()
	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to deleted user", err)
		return errors.NewInternalServerError("Database Error!")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := bookstores_users_db.Client.Prepare(QUERYFINDBYSTATUS)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = stmt.Close()
	}()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("Database Error!")
	}
	defer func() {
		err = rows.Close()
	}()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("Database Error!")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
