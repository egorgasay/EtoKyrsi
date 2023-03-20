package repository

import (
	"checkwork/internal/repository/prepared"
	"errors"
)

var WrongPasswordErr = errors.New("wrong password")

func (s Storage) CheckPassword(username string, password string) (bool, error) {
	stmt, err := prepared.GetPreparedStatement("CheckPassword")
	if err != nil {
		return false, err
	}

	var res int
	if stmt.QueryRow(username, password).Scan(&res) != nil {
		return false, err
	}

	if res < 1 {
		return false, WrongPasswordErr
	}

	return true, nil
}

//func (s Storage) CheckIsMentor(username string) (bool, error) {
//	stmt, err := prepared.GetPreparedStatement("CheckPassword")
//	if err != nil {
//		return false, err
//	}
//	var res int
//
//	if stmt.QueryRow(username).Scan(&res) != nil {
//		return false, err
//	}
//
//	if res < 1 {
//		return false, err
//	}
//
//	return true, nil
//}
