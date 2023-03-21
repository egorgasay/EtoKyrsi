package handlers

import (
	"checkwork/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) RegisterHandler(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get(config.Userkey)
	if user != nil {
		c.Redirect(http.StatusFound, "/logout")
		return
	}

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "reg.html", gin.H{})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")

	if password == "" && password2 == "" {
		c.HTML(http.StatusOK, "reg.html", gin.H{})
		return
	} else if password != password2 {
		c.HTML(http.StatusOK, "reg.html", gin.H{"err": "Passwords don't match"})
		return
	}

	err := h.logic.CreateUser(username, password)
	if err != nil {
		c.HTML(http.StatusOK, "reg.html", gin.H{"err": "Username is already taken"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, "/login")
}

func (h Handler) LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.Userkey)
	if user != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	status, err := h.logic.CheckPassword(username, password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"err": err})
		return
	}

	if !status && password != "" {
		c.HTML(http.StatusOK, "login.html", gin.H{"err": "Неверные данные!"})
		return
	} else if status && password != "" {
		session.Set(config.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"err": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusFound, "/")

		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h Handler) LoginMentorHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.Userkey)
	if user != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "mentorlogin.html", gin.H{})
		return
	}

	key := c.PostForm("password")

	if config.GetMentorKey() != key {
		c.HTML(http.StatusOK, "mentorlogin.html", gin.H{"err": "Неверные данные!"})
		return
	}

	session.Set(config.Userkey, key)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "mentorlogin.html", gin.H{"err": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/mentor")
}

func (h Handler) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete(config.Userkey)
	if err := session.Save(); err != nil {
		log.Println("Failed to delete session:", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func (h Handler) DocHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "task--100.htm", gin.H{})
}
