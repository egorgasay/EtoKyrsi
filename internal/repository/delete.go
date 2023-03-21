package repository

import "checkwork/internal/repository/prepared"

func (s Storage) DeletePullRequest(student string) error {
	stmt, err := prepared.GetPreparedStatement("DeletePullRequest")
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
	stmt, err := prepared.GetPreparedStatement("DeleteTask")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(num)
	if err != nil {
		return err
	}

	return nil
}
