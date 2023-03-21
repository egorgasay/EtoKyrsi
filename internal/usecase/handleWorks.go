package usecase

import "log"

func (uc *UseCase) SendPullRequest(line, student string) error {
	err := uc.storage.SetPending(student, 1)
	if err != nil {
		return err
	}

	err = uc.storage.AddPullRequest(line, student)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) HandleUserWork(username, student, verdict, msg string) error {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return NotAMentorError
	} else if !mentor {
		return NotAMentorError
	}

	err = uc.storage.DeletePullRequest(student)
	if err != nil {
		return err
	}

	err = uc.storage.SetPending(student, 0) // TODO: объеденить с SetVerdict
	if err != nil {
		return err
	}

	if verdict == "bad" {
		if msg == "" {
			msg = "В работе есть недочеты! Ментор оставил замечания на гитхабе."
		}

		err = uc.storage.SetVerdict(student, msg)
		if err != nil {
			return err
		}

		return nil
	}

	err = uc.storage.UpdateUserScore(student)
	if err != nil {
		return err
	}

	return nil
}
