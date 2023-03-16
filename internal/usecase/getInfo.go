package usecase

import (
	"bufio"
	"checkwork/internal/entity"
	"checkwork/internal/repository"
	"database/sql"
	"os"
	"strconv"
	"strings"
)

func (uc UseCase) GetWorks(username string) ([]repository.Work, error) {
	return uc.storage.GetWorks(username)
}

//func (uc UseCase) GetHTMLTask(number string) (func(obj any) (string, error), error) {
//	str, err := convertFileToString("templates/html/task-" + number + ".htm")
//	if err != nil {
//		return nil, err
//	}
//	var buf = make([]byte, 0)
//	buffer := bytes.NewBuffer(buf)
//	tmpl, err := template.New("task").Parse(str)
//	if err != nil {
//		return nil, err
//	}
//
//
//	return func(obj any) (string, error) {
//		err = tmpl.Execute(buffer, obj)
//		if err != nil {
//			return "", err
//		}
//		return buffer.String(), nil
//	}, err
//}

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

func (uc UseCase) GetTask(username string, number string) (task entity.Task, err error) {
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

func (uc UseCase) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	return uc.storage.GetTaskIDAndMsg(username)
}

func (uc UseCase) GetTasks(username string) ([]entity.Task, error) {
	return uc.storage.GetTasks(username)
}
