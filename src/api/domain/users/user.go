package users

import newrelic "github.com/newrelic/go-agent"

type User struct {
	Id int `json:"id" db:"id"`
	Username string `json:"username" db:"userName"`
}

type UsersService interface {
	FindAll(txn newrelic.Transaction) ([]User, error)
	Save(user User, txn newrelic.Transaction) error
}

type UsersRepository interface {
	FindAll() ([]User, error)
	Save(user User) error
}