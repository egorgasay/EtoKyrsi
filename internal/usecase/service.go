package usecase

import (
	"checkwork/internal/repository"
	"errors"
)

var NotAMentorError = errors.New("not a mentor")

type UseCase struct {
	storage repository.IStorage
}

func New(storage repository.IStorage) UseCase {
	if storage == nil {
		panic("storage is nil")
	}

	return UseCase{
		storage: storage,
	}
}
