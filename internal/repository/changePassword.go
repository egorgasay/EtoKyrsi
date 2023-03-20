package repository

import "checkwork/internal/repository/prepared"

func (s Storage) ChangePassword(username, oldPassword, newPassword string) error {
	_, err := s.CheckPassword(username, oldPassword)
	if err != nil {
		return err
	}

	stmt, err := prepared.GetPreparedStatement("ChangePassword")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newPassword, username)
	if err != nil {
		return err
	}

	return nil
}
