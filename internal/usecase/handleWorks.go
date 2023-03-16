package usecase

func (uc UseCase) SendPullRequest(line, student string) error {
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

func (uc UseCase) HandleUserWork(username, student, verdict string) error {
	err := uc.storage.DeletePullRequest(username, student)
	if err != nil {
		return err
	}

	err = uc.storage.SetPending(student, 0)
	if err != nil {
		return err
	}

	if verdict == "bad" {
		verdict = "В работе есть недочеты! Ментор оставил замечания на гитхабе."

		err := uc.storage.SetVerdict(student, verdict)
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
