package usecase

import (
	"checkwork/pkg/messageup"
	"os"
	"strconv"
)

func (uc UseCase) UpdateTasks(username, title, number, text string) error {
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
