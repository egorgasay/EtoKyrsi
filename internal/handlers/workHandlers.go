package handlers

import (
	"checkwork/internal/globals"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) SendWorkGetHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	username := user.(string)

	pending, err := h.logic.CheckIsPending(username)
	if err != nil {
		c.HTML(http.StatusOK, "send-work.html", gin.H{"error": err})
	}

	if pending {
		c.HTML(http.StatusOK, "404.html", gin.H{"sent": true})
		return
	}

	c.HTML(http.StatusOK, "send-work.html", gin.H{})
}

func (h Handler) SendWorkPostHandler(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get(globals.Userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	username := user.(string)

	pullURL := c.PostForm("pullURL")

	if pullURL == "" {
		c.Redirect(http.StatusFound, "/send")
		return
	}

	err := h.logic.SendPullRequest(pullURL, username)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "send-work.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "404.html", gin.H{"sent": true})
}
