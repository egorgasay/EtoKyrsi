package repository

import (
	"checkwork/internal/entity"
	"checkwork/internal/repository/prepared"
	"database/sql"
)

func (s Storage) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	stmt, err := prepared.GetPreparedStatement("GetTaskIDAndMsg")
	if err != nil {
		return 0, sql.NullString{}, err
	}
	var taskID int
	var msg sql.NullString
	var tasksCount int

	if err = stmt.QueryRow(username).Scan(&taskID, &msg, &tasksCount); err != nil {
		return 0, sql.NullString{}, err
	}

	if taskID > tasksCount-2 {
		taskID = 0
	}

	return taskID, msg, nil
}

type Work struct {
	Student string
	Link    string
}

func (s Storage) GetWorks(username string) ([]Work, error) {
	stmt, err := prepared.GetPreparedStatement("GetWorks")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
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
	stmt, err := prepared.GetPreparedStatement("GetTasks")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
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

func (s Storage) GetTitle(number int) (title string, err error) {
	stmt, err := prepared.GetPreparedStatement("GetTitle")
	if err != nil {
		return "", err
	}
	return title, stmt.QueryRow(number).Scan(&title)
}

func (s Storage) GetUsers() ([]entity.User, error) {
	stmt, err := prepared.GetPreparedStatement("GetUsers")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	var users = make([]entity.User, 0)
	for rows.Next() {
		var user entity.User

		err = rows.Scan(&user.Name, &user.Level, &user.LastComment)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}
