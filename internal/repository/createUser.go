package repository

import (
	"checkwork/internal/repository/prepared"
	"errors"
	"github.com/mattn/go-sqlite3"
)

var AlreadyExistsErr = errors.New("already exists")

func (s Storage) CreateUser(username string, password string) error {
	stmt, err := prepared.GetPreparedStatement("CreateUser")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, password)

	if errors.Is(err, sqlite3.ErrConstraintUnique) {
		return AlreadyExistsErr
	}

	return err
}
