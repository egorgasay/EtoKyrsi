package repository

import (
	"errors"
	"github.com/mattn/go-sqlite3"
)

var AlreadyExistsErr = errors.New("already exists")

func (s Storage) CreateUser(username string, password string) error {
	query := "INSERT INTO Users (student, password, task_id, pending) VALUES (?, ?, 1, 0)"
	_, err := s.DB.Exec(query, username, password)

	if errors.Is(err, sqlite3.ErrConstraintUnique) {
		return AlreadyExistsErr
	}

	return err
}
