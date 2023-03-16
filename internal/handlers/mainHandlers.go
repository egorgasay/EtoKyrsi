package handlers

import (
	"checkwork/internal/globals"
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
	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	username, _ := user.(string)

	id, msg, err := h.logic.GetTaskIDAndMsg(username)
	if err != nil {
		// TODO:
		c.HTML(http.StatusOK, "task.htm", gin.H{"error": err})
		return
	}

	filename := fmt.Sprintf("task-%d.htm", id)

	if msg.String != "" {
		c.HTML(http.StatusOK, filename, gin.H{"error": msg.String,
			"IsPending": false})
		return
	}

	pending, err := h.logic.CheckIsPending(username)
	if err != nil {
		c.HTML(http.StatusOK, filename, gin.H{"error": err})
		return
	}

	pullURL := c.PostForm("pullURL")
	if pullURL != "" {
		err := h.logic.SendPullRequest(pullURL, username)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusOK, filename, gin.H{"error": err, "Username": username})
			return
		}
		pending = true
	}

	c.HTML(http.StatusOK, filename, gin.H{"IsPending": pending, "Username": username})
}

func (h Handler) NotFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{"page": "404"})
}
