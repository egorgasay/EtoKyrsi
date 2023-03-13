package usecase

import "errors"

var ErrAlreadyExists = errors.New("username already taken")

func (uc *UseCase) CreateUser(username, password string) error {
	isMentor, err := uc.storage.CheckIsMentor(username)
	if err != nil {
		return err
	}

	if isMentor {
		return ErrAlreadyExists
	}

	return uc.storage.CreateUser(username, password)
}

func (uc *UseCase) CheckPassword(username, password string) (bool, error) {
	return uc.storage.CheckPassword(username, password)
}
