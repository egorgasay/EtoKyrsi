package repository

func (s Storage) DeletePullRequest(username, student string) error {
	query := `
	DELETE FROM PullRequests WHERE student = ?
`
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

func (s Storage) DeleteTask(num int) error {
	query := `
	DELETE FROM Tasks WHERE task_id = ?
`
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(num)
	if err != nil {
		return err
	}

	return nil
}
