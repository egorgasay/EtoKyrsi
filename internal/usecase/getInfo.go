package usecase

import (
	"bufio"
	"checkwork/internal/entity"
	"checkwork/internal/repository"
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
)

func (uc *UseCase) GetWorks(username string) ([]repository.Work, error) {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return nil, NotAMentorError
	} else if !mentor {
		return nil, NotAMentorError
	}

	return uc.storage.GetWorks()
}

func (uc *UseCase) GetUsers(username string) ([]entity.User, error) {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return nil, NotAMentorError
	} else if !mentor {
		return nil, NotAMentorError
	}

	return uc.storage.GetUsers()
}

func convertFileToString(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0777)
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

func (uc *UseCase) GetTask(username string, number string) (task entity.Task, err error) {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return entity.Task{}, NotAMentorError
	} else if !mentor {
		return entity.Task{}, NotAMentorError
	}

	str, err := convertFileToString("templates/mup/task-" + number + ".mup")
	if err != nil {
		return entity.Task{}, err
	}

	num, err := strconv.Atoi(number)
	if err != nil {
		return entity.Task{}, err
	}

	title, err := uc.storage.GetTitle(num)
	if err != nil {
		return entity.Task{}, err
	}

	task.Name = title
	task.Number = num
	task.Text = str

	return task, nil
}

func (uc *UseCase) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	return uc.storage.GetTaskIDAndMsg(username)
}

func (uc *UseCase) GetTasks(username string) ([]entity.Task, error) {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return nil, NotAMentorError
	} else if !mentor {
		return nil, NotAMentorError
	}

	return uc.storage.GetTasks()
}
