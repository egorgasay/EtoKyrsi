package usecase

import (
	"checkwork/pkg/messageup"
	"log"
	"os"
	"strconv"
)

func (uc *UseCase) UpdateTasks(username, title, number, text string) error {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return NotAMentorError
	} else if !mentor {
		return NotAMentorError
	}

	file, err := os.OpenFile("templates/mup/task-"+number+".mup", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(text))
	if err != nil {
		return err
	}
	file.Close()

	html, err := messageup.ToHTML(text)

	file, err = os.OpenFile("templates/html/task-"+number+".htm", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(html))
	if err != nil {
		return err
	}

	num, err := strconv.Atoi(number)
	if err != nil {
		return err
	}

	err = uc.storage.UpdateTask(num, title)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteTasks(username, number string) error {
	mentor, err := uc.CheckIsMentor(username) // TODO: REPLACE WITH MIDDLEWARE
	if err != nil {
		log.Println(err)
		return NotAMentorError
	} else if !mentor {
		return NotAMentorError
	}

	num, err := strconv.Atoi(number)
	if err != nil {
		return err
	}

	return uc.storage.DeleteTask(num)
}
