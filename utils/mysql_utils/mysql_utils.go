package mysql_utils

import (
	"bookstoreUsersApi/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ERRORNOROWS = "no rows in result set"
)

func ParsError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ERRORNOROWS) {
			return errors.NewNotFoundError("No record match given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processin request")
}
