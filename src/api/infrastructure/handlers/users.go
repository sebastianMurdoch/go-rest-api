package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users"
	"net/http"
)

type UsersHandler struct {
	S  users.UsersService `inject:"auto"`
	Nr newrelic.Application `inject:"auto"`
}

func (u *UsersHandler) Get(c *gin.Context) {
	var txn newrelic.Transaction

	// Used for localhost tests
	if u.Nr != nil {
		txn = u.Nr.StartTransaction("GET /users", c.Writer, c.Request)
		defer txn.End()
	}

	users, err := u.S.FindAll(txn)
	if err != nil {
		txn.NoticeError(err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, users)
}

func (u *UsersHandler) Post(c *gin.Context) {
	var txn newrelic.Transaction
	if u.Nr != nil {
		txn = u.Nr.StartTransaction("POST /users", c.Writer, c.Request)
		defer txn.End()
	}

	var user users.User
	err := c.Bind(&user)
	if err != nil {
		if txn != nil {
			txn.NoticeError(err)
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = u.S.Save(user, txn)
	if err != nil {
		newErr := errors.New("Error at UsersHandler-Post -- "+err.Error())
		if txn != nil {
			txn.NoticeError(newErr)
		}
		c.AbortWithError(http.StatusInternalServerError, newErr)
	}

	c.JSON(http.StatusCreated, nil)
}
