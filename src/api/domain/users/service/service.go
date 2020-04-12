package service

import (
	"errors"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users"
)

type UserServiceImpl struct {
	Repository users.UsersRepository
}

func NewUserServiceImpl(repository users.UsersRepository) *UserServiceImpl  {
	return &UserServiceImpl{
		Repository: repository,
	}
}

func (u *UserServiceImpl) FindAll(txn newrelic.Transaction) ([]users.User, error)  {
	seg := newrelic.StartSegment(txn, "UserServiceImpl-FindAll")
	defer seg.End()
	users, err := u.Repository.FindAll()
	if err != nil {
		newErr := errors.New("Error at UserServiceImpl-FindAll -- "+err.Error())
		txn.NoticeError(newErr)
		return nil, newErr
	}
	return users, nil
}

func (u *UserServiceImpl) Save(user users.User, txn newrelic.Transaction) error {
	seg := newrelic.StartSegment(txn, "UserServiceImpl-Save")
	defer seg.End()
	err := u.Repository.Save(user)
	if err != nil {
		newError := errors.New("Error at UserServiceImpl-Save -- "+err.Error())
		txn.NoticeError(newError)
		return newError
	}
	return nil
}

