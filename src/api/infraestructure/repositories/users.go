package repositories

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/users"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl{
	return &UserRepositoryImpl{DB:db}
}

func (r *UserRepositoryImpl) FindAll() ([]users.User, error) {
	allUsers := []users.User{}
	err := r.DB.Select(&allUsers, "SELECT * FROM users")
	if err != nil {
		return nil, errors.New("Error at UserRepositoryImpl-FindAll -- "+err.Error())
	}
	return allUsers, nil
}

func (r *UserRepositoryImpl) Save(user users.User) error {
	insertQuery := `INSERT INTO users (userName) VALUES (?)`
	_, err := r.DB.Exec(insertQuery,user.Username)
	if err != nil {
		return errors.New("Error at UserRepositoryImpl-Save -- "+err.Error())
	}
	return nil
}