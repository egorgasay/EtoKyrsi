package handlers

import (
	"checkwork/internal/entity"
	"checkwork/internal/globals"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h Handler) MentorGetHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username, _ := user.(string)

	isMentor, err := h.logic.CheckIsMentor(username)
	if err != nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": "В доступе отказано"})
		return
	}

	if !isMentor {
		c.Redirect(http.StatusFound, "/")
		return
	}

	works, err := h.logic.GetWorks(username)
	if err != nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": err})
		return
	}

	lessons, err := h.logic.GetTasks(username)
	if err != nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{"Works": works, "Lessons": lessons, "Count": len(lessons)})
}

func (h Handler) MentorPostHandler(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username := user.(string)
	pullURL := c.PostForm("verdict")
	if pullURL == "" {
		c.Redirect(http.StatusFound, "/mentor")
		return
	}

	split := strings.Split(pullURL, " - ")

	if len(split) != 2 {
		log.Println("Неправильный вердикт")
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": "Ошибка сервера"})
		return
	}

	student, status := split[0], split[1]
	err := h.logic.HandleUserWork(username, student, status)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": err})
		return
	}

	works, err := h.logic.GetWorks(username)
	if err != nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{"Works": works})
}

func (h Handler) ViewTask(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username, _ := user.(string)

	number := c.Request.URL.Query().Get("select")
	h.viewHelper(c, username, number)
}

func (h Handler) viewHelper(c *gin.Context, username, number string) {
	var task entity.Task
	if number == "NEW" {
		count := c.PostForm("count")
		if count == "" {
			count = "0"
		}
		atoi, err := strconv.Atoi(count)
		if err != nil {
			return
		}

		task.Number = atoi + 1
		c.HTML(http.StatusOK, "change-task.html", gin.H{"Task": task})
		return
	}

	task, err := h.logic.GetTask(username, number)
	if err != nil {
		c.HTML(http.StatusOK, "change-task.html", gin.H{"error": err})
		return
	}
	name := fmt.Sprintf("task-%d.htm", task.Number)
	c.HTML(http.StatusOK, name, gin.H{"Task": task})
}

func (h Handler) ChangeTask(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username, _ := user.(string)

	number := c.PostForm("number")
	title := c.PostForm("title")
	text := c.PostForm("text")

	if err := h.logic.UpdateTasks(username, title, number, text); err != nil {
		c.HTML(http.StatusOK, "change-task.html", gin.H{"error": err})
		return
	}

	h.viewHelper(c, username, number)
}
