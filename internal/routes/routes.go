package routes

import (
	"checkwork/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.RouterGroup, h handlers.Handler) {
	r.Any("/reg", h.RegisterHandler)
	r.Any("/login", h.LoginHandler)
	r.Any("/mentor-login", h.LoginMentorHandler)
	r.Any("/logout", h.LogoutHandler)
}

func PrivateRoutes(r *gin.RouterGroup, h handlers.Handler) {
	r.Any("/", h.MainHandler)

	r.GET("/mentor", h.MentorGetHandler)
	r.POST("/mentor", h.MentorPostHandler)

	r.GET("/send", h.SendWorkGetHandler)
	r.POST("/send", h.SendWorkPostHandler)
}
