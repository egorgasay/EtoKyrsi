package handlers

import (
	"checkwork/config"
	"checkwork/internal/usecase"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	logic usecase.UseCase
}

func NewHandler(logic usecase.UseCase) *Handler {
	return &Handler{logic: logic}
}

func (h Handler) MainHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username, _ := user.(string)

	id, msg, err := h.logic.GetTaskIDAndMsg(username)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "task--1.htm", gin.H{})
		return
	}

	filename := fmt.Sprintf("task-%d.htm", id)

	var pending = false
	pullURL := c.PostForm("pullURL")
	if pullURL != "" {
		err = h.logic.SendPullRequest(pullURL, username)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "task--1.htm", gin.H{})
			return
		}
		pending = true
	} else {
		pending, err = h.logic.CheckIsPending(username)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, "task--1.htm", gin.H{})
			return
		}

		if msg.String != "" && !pending {
			c.HTML(http.StatusOK, filename, gin.H{"error": msg.String,
				"IsPending": false})
			return
		}
	}

	c.HTML(http.StatusOK, filename, gin.H{"IsPending": pending, "Username": username})
}

func (h Handler) NotFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{"page": "404"})
}
