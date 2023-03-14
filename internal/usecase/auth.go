package usecase

import "errors"

var ErrAlreadyExists = errors.New("username already taken")

func (uc *UseCase) CreateUser(username, password string) error {
	return uc.storage.CreateUser(username, password)
}

func (uc *UseCase) CheckPassword(username, password string) (bool, error) {
	return uc.storage.CheckPassword(username, password)
}
