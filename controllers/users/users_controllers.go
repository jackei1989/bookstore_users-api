package users

import (
	"bookstoreUsersApi/domain/users"
	"bookstoreUsersApi/services"
	"bookstoreUsersApi/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented!")
}
