package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users"
	"net/http"
)

type UsersHandler struct {
	s users.UsersService
	nr newrelic.Application
}

func NewUsersHandler(s users.UsersService, nr newrelic.Application) *UsersHandler{
	return &UsersHandler{s:s, nr:nr}
}

func (u *UsersHandler) Get(c *gin.Context) {
	txn := u.nr.StartTransaction("GET /users", c.Writer, c.Request)
	defer txn.End()

	users, err := u.s.FindAll(txn)
	if err != nil {
		txn.NoticeError(err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, users)
}

func (u *UsersHandler) Post(c *gin.Context) {
	txn := u.nr.StartTransaction("POST /users", c.Writer, c.Request)
	defer txn.End()

	var user users.User
	err := c.Bind(&user)
	if err != nil {
		txn.NoticeError(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = u.s.Save(user, txn)
	if err != nil {
		newErr := errors.New("Error at UsersHandler-Post -- "+err.Error())
		txn.NoticeError(newErr)
		c.AbortWithError(http.StatusInternalServerError, newErr)
	}

	c.JSON(http.StatusCreated, nil)
}
