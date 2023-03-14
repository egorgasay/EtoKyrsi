package repository

func (s Storage) SetPending(username string, status int) error {
	_, err := s.DB.Exec(
		"UPDATE Users SET pending = $1 WHERE student = $2",
		status, username)
	return err
}

func (s Storage) CheckIsPending(username string) (bool, error) {
	query := "SELECT pending FROM Users WHERE student = $1"
	row := s.DB.QueryRow(query, username)

	err := row.Err()
	if err != nil {
		return false, err
	}

	var isPending int
	err = row.Scan(&isPending)
	if err != nil {
		return false, err
	}

	return 1 == isPending, nil
}
