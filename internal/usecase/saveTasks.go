package usecase

import "os"

func (uc UseCase) UpdateHtml(username string, html string) error {
	file, err := os.OpenFile("templates/html/task.htm", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(html))
	if err != nil {
		return err
	}

	return nil
}
