package repository

func (s Storage) DeletePullRequest(username, student string) error {
	query := `
	DELETE FROM PullRequests WHERE student = $1
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
