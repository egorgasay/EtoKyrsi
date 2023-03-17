package repository

import "log"

func (s Storage) SetVerdict(student, verdict string) error {
	query := `
	UPDATE Users 
	SET msg = ?
	WHERE student = ?;
`
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(verdict, student)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) UpdateUserScore(student string) error {
	query := `
UPDATE Users 
SET 
    task_id = task_id + 1, 
	msg = ''
WHERE student = ?`

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(student)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) UpdateTask(num int, title string) error {
	query := `INSERT OR REPLACE INTO Tasks (task_id, title) VALUES (?, ?)`

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(num, title)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s Storage) AddPullRequest(link, student string) error {
	query := `INSERT INTO PullRequests (link, student) VALUES (?, ?)`

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(link, student)
	if err != nil {
		return err
	}

	return nil
}
