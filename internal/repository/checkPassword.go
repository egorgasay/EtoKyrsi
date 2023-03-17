package repository

func (s Storage) CheckPassword(username string, password string) (bool, error) {
	query := "SELECT count(*) FROM Users WHERE student = ? AND password = ?"
	var res int

	row := s.DB.QueryRow(query, username, password)
	err := row.Scan(&res)

	if err != nil {
		return false, err
	}

	if res < 1 {
		return false, err
	}

	return true, nil
}

//func (s Storage) CheckMentorPassword(username string, password string) (bool, error) {
//	query := "SELECT count(*) FROM Mentors WHERE username = ? AND password = ?"
//	var res int
//
//	row := s.DB.QueryRow(query, username, password)
//	err := row.Scan(&res)
//	if err != nil {
//		return false, err
//	}
//
//	if res < 1 {
//		return false, err
//	}
//
//	return true, nil
//}

func (s Storage) CheckIsMentor(username string) (bool, error) {
	query := "SELECT count(*) FROM Mentors WHERE username = ?"
	var res int

	row := s.DB.QueryRow(query, username)
	err := row.Scan(&res)
	if err != nil {
		return false, err
	}

	if res < 1 {
		return false, err
	}

	return true, nil
}
