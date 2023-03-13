package main

import (
	"checkwork/config"
	"checkwork/internal/globals"
	"checkwork/internal/handlers"
	"checkwork/internal/middleware"
	"checkwork/internal/repository"
	"checkwork/internal/routes"
	"checkwork/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()
	cfg := config.New()

	storage, err := repository.Init(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	logic := usecase.New(storage)

	h := handlers.NewHandler(logic)

	r.LoadHTMLGlob("templates/html/*")
	r.Static("/static", "static")
	r.NoRoute(h.NotFoundHandler)

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	routes.PublicRoutes(public, *h)

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, *h)

	go r.Run(cfg.Host)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown market ...")
}
