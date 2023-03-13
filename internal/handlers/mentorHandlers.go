package handlers

import (
	"checkwork/internal/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	tasks, err := h.logic.GetTasks(username)
	if err != nil {
		c.HTML(http.StatusOK, "admin.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{"Works": works, "Tasks": tasks})
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
	html := c.PostForm("html")
	if pullURL == "" && html != "" {
		err := h.logic.UpdateHtml(username, html)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "admin.html", gin.H{"error": "Ошибка сервера"})
		}
		c.Redirect(http.StatusFound, "/mentor")
		return
	} else if pullURL == "" {
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
