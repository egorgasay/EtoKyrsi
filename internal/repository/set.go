package repository

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

func (s Storage) UpdateTask(student string) error {
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
