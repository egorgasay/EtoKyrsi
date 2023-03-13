package repository

func (s Storage) SetPending(username string, status int) error {
	_, err := s.DB.Exec(
		"UPDATE Users SET pending = ? WHERE student = ?",
		status, username)
	return err
}

func (s Storage) CheckIsPending(username string) (bool, error) {
	query := "SELECT pending FROM Users WHERE student = ?"
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
