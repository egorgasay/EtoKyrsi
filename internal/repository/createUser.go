package repository

func (s Storage) CreateUser(username string, password string) error {
	query := "INSERT INTO Users (student, password, task_id, pending) VALUES (?, ?, 1, 0)"
	_, err := s.DB.Exec(query, username, password)
	return err
}
