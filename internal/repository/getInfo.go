package repository

import (
	"checkwork/internal/entity"
	"database/sql"
)

func (s Storage) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	query := "SELECT task_id, msg FROM Users WHERE student = $1"
	row := s.DB.QueryRow(query, username)

	err := row.Err()
	if err != nil {
		return 0, sql.NullString{}, err
	}

	var taskID int
	var msg sql.NullString
	err = row.Scan(&taskID, &msg)
	if err != nil {
		return 0, sql.NullString{}, err
	}

	return taskID, msg, nil
}

type Work struct {
	Student string
	Link    string
}

func (s Storage) GetWorks(username string) ([]Work, error) {
	query := "SELECT student, link FROM PullRequests"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	var works = make([]Work, 0, 10)
	for rows.Next() {
		var work Work
		err = rows.Scan(&work.Student, &work.Link)
		if err != nil {
			return nil, err
		}

		works = append(works, work)
	}

	return works, nil
}

func (s Storage) GetTasks(username string) ([]entity.Task, error) {
	query := "SELECT task_id, title FROM Tasks"
	prepare, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := prepare.Query()
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	var tasks = make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task

		err = rows.Scan(&task.Number, &task.Name)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s Storage) GetTitle(number int) (string, error) {
	query := "SELECT title FROM Tasks WHERE task_id = $1"
	prepare, err := s.DB.Prepare(query)
	if err != nil {
		return "", err
	}

	row := prepare.QueryRow(number)
	err = row.Err()
	if err != nil {
		return "", err
	}

	var title string
	return title, row.Scan(&title)
}
