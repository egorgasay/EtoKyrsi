package messageup

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
)

var begin Template = `<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      href="https://fonts.googleapis.com/css2?family=Fira+Code:wght@600&family=Ubuntu:wght@400;500;700&display=swap"
      rel="stylesheet"
    />
    <title>etocursi</title>
    <link rel="stylesheet" href="/static/css/reset.css" />
    <link rel="stylesheet" href="/static/css/task.css" />
    <link rel="stylesheet" href="/static/css/message.css">
  </head>
  <body>
	<header>
      <div class="exit">
        <a href="/logout"><img src="/static/img/exit.svg" alt="exit"/></a>
      </div>
    </header>
`

var msg Template = `<div class="text">
	{{ . }}
</div>
`

var msgWarn Template = `<div class="text important">
	Важно! Не мержите код в основную ветку до того как я приму код.
</div>
`

var inputConfirm Template = ` {{ if and .task (not .IsPending) }}
    <br>
      <form method="post">
        <center>
        <input name="pullURL" class="input" placeholder="Ссылка на решение" required>
      <button type="submit" class="btn">Отправить решение</button>
      </center>
      </form>
      {{end}}
    {{ if .IsPending}}
    <br>
    <center>
      <p>Ментор еще не проверил вашу работу!</p>
    </center>
    {{ end }}`

var lesson Template = `<section class="lesson">
	<div class="container">
`

var endLesson Template = `	</div>
</section>
`

var header Template = `<div class="start">
	<h1 class="h1">{{ . }}</h1>
</div>
`

var newLine Template = `<br>`

var msgHeader Template = `<h2 class="h2">{{ . }}</h2>`

var end Template = `
    <script src="/static/js/message.js"></script>
  </body>
</html>
`

var buttonALink Template = `<div class="start-btn">
	<a href="{{.Link}}"
              target="_blank"
              class="start-btn__link">{{.Text}}</a>
</div>
`

var simpleText Template = `<div class="simple-text">
	{{.}}
</div>`

type CommandAndText struct {
	Command action
	Text    []string
}

type action string
type Template string
type Middleware func(lines []string) (string, error)

var actions = map[action]Middleware{
	"@msg":           doMSG,
	"@header":        doHeader,
	"@lesson":        doLesson,
	"@endlesson":     doEndLesson,
	"@msg-warn":      doWarn,
	"@confirm-input": doInputConfirm,
	"@msg-header":    doMSGHeader,
	"@n":             doNewLine,
	"@btn-link":      doButtonLink,
	"@st":            doSimpleText,
}

func doMSG(lines []string) (string, error) {
	return do(lines, msg)
}

func doNewLine(lines []string) (string, error) {
	return do(lines, newLine)
}

func doWarn(lines []string) (string, error) {
	return do(lines, msgWarn)
}

func doMSGHeader(lines []string) (string, error) {
	return do(lines, msgHeader)
}

func doHeader(lines []string) (string, error) {
	return do(lines, header)
}

func doLesson(lines []string) (string, error) {
	return do(lines, lesson)
}
func doEndLesson(lines []string) (string, error) {
	return do(lines, endLesson)
}

func doInputConfirm(lines []string) (string, error) {
	return string(inputConfirm), nil
}

func doSimpleText(lines []string) (string, error) {
	return do(lines, simpleText)
}

type buttonLink struct {
	Text string
	Link string
}

var ErrWrongButton = errors.New("wrong button declaration")

func doButtonLink(lines []string) (string, error) {
	var bl buttonLink
	lns := strings.Join(lines, "!")
	btn := strings.Split(lns, "!")
	if len(btn) != 4 {
		return "", ErrWrongButton
	}

	bl.Text = btn[1]
	bl.Link = btn[3]

	tmpl, err := template.New("tmpl").Parse(string(buttonALink))
	if err != nil {
		return "", err
	}

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err = tmpl.Execute(buffer, bl)
	if err != nil {
		return "", err
	}
	return buffer.String(), err
}

func do(lines []string, Template Template) (string, error) {
	tmpl, err := template.New("tmpl").Parse(string(Template))
	if err != nil {
		return "", err
	}
	text := strings.Join(lines, "\n")

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err = tmpl.Execute(buffer, text)
	if err != nil {
		return "", err
	}
	return buffer.String(), err
}

func ToHTML(mup string) (string, error) {
	lines := strings.Split(mup, "\r\n")
	var storage = make([]string, 0)
	br := false
	var cat CommandAndText

	storage = append(storage, string(begin))
	for _, line := range lines {
		if len(line) < 1 || line == "\n" {
			br = true
			act, ok := actions[cat.Command]
			if !ok {
				continue
			}

			txt, err := act(cat.Text)
			if err != nil {
				return "", err
			}

			storage = append(storage, txt)
			cat = CommandAndText{}
			continue
		}

		if line[0] == '@' {
			br = false
			cmd := strings.Split(line, " ")
			cat.Command = action(cmd[0])
			continue
		}

		if !br {
			cat.Text = append(cat.Text, line)
		} else {
			cat.Command = action("@st")
			cat.Text = append(cat.Text, line)
		}
	}
	storage = append(storage, string(end))

	return strings.Join(storage, "\n"), nil
}
