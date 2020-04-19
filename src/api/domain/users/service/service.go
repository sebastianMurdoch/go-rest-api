package service

import (
	"errors"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users"
)

type UserServiceImpl struct {
	R users.UsersRepository `inject:"auto"`
}

func (u *UserServiceImpl) FindAll(txn newrelic.Transaction) ([]users.User, error)  {
	if txn != nil {
		seg := newrelic.StartSegment(txn, "UserServiceImpl-FindAll")
		defer seg.End()
	}
	users, err := u.R.FindAll()
	if err != nil {
		newErr := errors.New("Error at UserServiceImpl-FindAll -- "+err.Error())
		if txn != nil {
			txn.NoticeError(newErr)
		}
		return nil, newErr
	}
	return users, nil
}

func (u *UserServiceImpl) Save(user users.User, txn newrelic.Transaction) error {
	if txn != nil {
		seg := newrelic.StartSegment(txn, "UserServiceImpl-Save")
		defer seg.End()
	}
	err := u.R.Save(user)
	if err != nil {
		newError := errors.New("Error at UserServiceImpl-Save -- "+err.Error())
		if txn != nil {
			txn.NoticeError(newError)
		}
		return newError
	}
	return nil
}

