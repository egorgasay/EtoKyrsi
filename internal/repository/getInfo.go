package repository

import "database/sql"

func (s Storage) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	query := "SELECT task_id, msg FROM Users WHERE student = ?"
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
