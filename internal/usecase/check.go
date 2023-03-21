package usecase

import (
	"checkwork/config"
)

func (uc *UseCase) CheckIsPending(username string) (bool, error) {
	return uc.storage.CheckIsPending(username)
}

func (uc *UseCase) CheckIsMentor(username string) (bool, error) {
	return username == config.GetMentorKey(), nil
}
