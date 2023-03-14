package usecase

import (
	"bufio"
	"checkwork/internal/repository"
	"database/sql"
	"os"
	"strings"
)

func (uc UseCase) GetWorks(username string) ([]repository.Work, error) {
	return uc.storage.GetWorks(username)
}

func (uc UseCase) GetTasks(username string) (string, error) {
	file, err := os.OpenFile("templates/html/tasks-mup.htm", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var sb = strings.Builder{}

	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}

	return sb.String(), nil
}

func (uc UseCase) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	return uc.storage.GetTaskIDAndMsg(username)
}
