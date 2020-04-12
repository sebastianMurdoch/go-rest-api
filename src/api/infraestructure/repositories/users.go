package repositories

import (
	"database/sql"
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
	var id int
	getNextID := `SELECT id FROM users ORDER BY id DESC LIMIT 1`
	row := r.DB.QueryRow(getNextID)
	err := row.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("Error at UserRepositoryImpl-Save -- "+ err.Error())
	}
	if id == 0 {
		id = 1
	} else {
		id += 1
	}

	insertQuery := `INSERT INTO users(id, userName) VALUES (?, ?)`
	_, err = r.DB.Exec(insertQuery, id, user.Username)
	if err != nil {
		return errors.New("Error at UserRepositoryImpl-Save -- "+err.Error())
	}
	return nil
}