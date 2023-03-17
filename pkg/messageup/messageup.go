package messageup

import (
	"bytes"
	"errors"
	"strconv"
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
{{ template "change" .}}
	<header>
	{{ if .error }}
    {{ template "msg" .}}
    {{ end }}
      <div class="exit">
        <a href="/logout"><img src="/static/img/exit.svg" alt="exit"/></a>
      </div>
    </header>
	<div class="container">
`

var msg Template = `<div class="text">
	{{ . }}
</div>
`

var msgWarn Template = `<div class="text important">
	{{ . }}
</div>
`

var lessonLogic Template = `if or (eq .task {{ . }} ) (not .task)`

var lesson Template = `
	<section class="lesson">
		<div class="container">
`

var endLesson Template = `{{ if and (not .Task) (not .IsPending) }}
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
    {{ end }}
	</div>
</section>
`

var header Template = `<div class="start">
	<h1 class="h1">{{ . }}</h1>
</div>
`

var newLine Template = `<br>`

var msgHeader Template = `<h2 class="h2">{{ . }}</h2>`

var text Template = `<h3 style="font-size:32px">{{ . }}</h3>`

var end Template = endLesson + `
	</div>
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
	Params  []string
	Text    []string
}

type action string
type Template string
type Middleware func(lines string, params []string) (string, error)

var actions = map[action]Middleware{
	"@msg":        doMSG,
	"@header":     doHeader,
	"@lesson":     doLesson,
	"@endlesson":  doEndLesson,
	"@msg-warn":   doWarn,
	"@msg-header": doMSGHeader,
	"@n":          doNewLine,
	"@btn-link":   doButtonLink,
	"@st":         doSimpleText,
	"@text":       doText,
}

func doMSG(lines string, params []string) (string, error) {
	return do(lines, msg)
}

func doText(lines string, params []string) (string, error) {
	return do(lines, text)
}

func doNewLine(lines string, params []string) (string, error) {
	return do(lines, newLine)
}

func doWarn(lines string, params []string) (string, error) {
	return do(lines, msgWarn)
}

func doMSGHeader(lines string, params []string) (string, error) {
	return do(lines, msgHeader)
}

func doHeader(lines string, params []string) (string, error) {
	return do(lines, header)
}

var WrongNumberOfLesson = errors.New("wrong number of lesson")

func doLesson(lines string, params []string) (string, error) {
	if params == nil || len(params) == 0 {
		return "", WrongNumberOfLesson
	}

	num, err := strconv.Atoi(params[0])
	if err != nil {
		return "", WrongNumberOfLesson
	}

	res, err := do(num, lessonLogic)
	if err != nil {
		return "", err
	}

	res = "{{" + res + "}}"
	return res + string(lesson), nil
}

func doEndLesson(lines string, params []string) (string, error) {
	return string(endLesson), nil
}

func doSimpleText(lines string, params []string) (string, error) {
	return do(lines, simpleText)
}

type buttonLink struct {
	Text string
	Link string
}

var ErrWrongButton = errors.New("wrong button declaration")

func doButtonLink(lines string, params []string) (string, error) {
	var bl buttonLink
	lns := strings.Split(lines, "\n")

	if len(lns) != 2 {
		return "", ErrWrongButton
	}

	text := strings.Split(lns[0], "!")
	if len(text) != 2 {
		return "", ErrWrongButton
	}

	link := strings.Split(lns[1], "!")
	if len(link) != 2 {
		return "", ErrWrongButton
	}

	bl.Text = text[1]
	bl.Link = link[1]

	return do(bl, buttonALink)
}

func do(obj any, Template Template) (string, error) {
	tmpl, err := template.New("tmpl").Parse(string(Template))
	if err != nil {
		return "", err
	}

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err = tmpl.Execute(buffer, obj)
	if err != nil {
		return "", err
	}
	return buffer.String(), err
}

func ToHTML(mup string) (string, error) {
	lines := strings.Split(mup+"\r\n\n\r\n\n\r\n@n", "\r\n")
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

			text := strings.Join(cat.Text, "\n")
			txt, err := act(text, cat.Params)
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

			if len(cmd) > 1 {
				cat.Params = strings.Split(cmd[1], " ")
			}
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
