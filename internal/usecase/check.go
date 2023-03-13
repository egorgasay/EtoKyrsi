package usecase

import "checkwork/internal/globals"

func (uc UseCase) CheckIsPending(username string) (bool, error) {
	return uc.storage.CheckIsPending(username)
}

func (uc UseCase) CheckIsMentor(username string) (bool, error) {
	return username == string(globals.Secret), nil
}
