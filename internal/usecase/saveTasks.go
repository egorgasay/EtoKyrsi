package usecase

import (
	"checkwork/pkg/messageup"
	"os"
)

func (uc UseCase) UpdateTasks(username string, mup string) error {
	file, err := os.OpenFile("templates/html/tasks-mup.htm", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(mup))
	if err != nil {
		return err
	}
	file.Close()

	html, err := messageup.ToHTML(mup)

	file, err = os.OpenFile("templates/html/task.htm", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(`{{define "tasks"}}` + html + `{{end}}`))
	if err != nil {
		return err
	}

	return nil
}
