package repository

import (
	"checkwork/internal/repository/prepared"
	"log"
)

func (s Storage) SetVerdict(student, verdict string) error {
	stmt, err := prepared.GetPreparedStatement("SetVerdict")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(verdict, student)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) UpdateUserScore(student string) error {
	stmt, err := prepared.GetPreparedStatement("UpdateUserScore")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(student)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) UpdateTask(num int, title string) error {
	stmt, err := prepared.GetPreparedStatement("UpdateTask")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(num, title)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s Storage) AddPullRequest(link, student string) error {
	stmt, err := prepared.GetPreparedStatement("AddPullRequest")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(link, student)
	if err != nil {
		return err
	}

	return nil
}
