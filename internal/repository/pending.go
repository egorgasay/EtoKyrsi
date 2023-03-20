package repository

import "checkwork/internal/repository/prepared"

func (s Storage) SetPending(username string, status int) error {
	stmt, err := prepared.GetPreparedStatement("SetPending")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, username)
	return err
}

func (s Storage) CheckIsPending(username string) (bool, error) {
	stmt, err := prepared.GetPreparedStatement("CheckIsPending")
	if err != nil {
		return false, err
	}

	var isPending int
	err = stmt.QueryRow(username).Scan(&isPending)
	if err != nil {
		return false, err
	}

	return 1 == isPending, nil
}
