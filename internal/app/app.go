package app

import (
	"log/slog"

	"github.com/avran02/medods/config"
	"github.com/avran02/medods/internal/controller"
	"github.com/avran02/medods/internal/pkg/jwt"
	"github.com/avran02/medods/internal/repository"
	"github.com/avran02/medods/internal/router"
	"github.com/avran02/medods/internal/service"
	"github.com/avran02/medods/logger"
)

type App struct {
	router     *router.Router
	repository repository.Repository
}

func New() *App {
	config := config.New()
	logger.Setup(config.ServerConfig)
	repo := repository.New(&config.DBConfig)
	jwtGenerator := jwt.NewJwtGenerator(config.JWTConfig)
	service := service.New(repo, jwtGenerator, config.SMTPConfig)
	controller := controller.New(service)
	router := router.New(controller, config.ServerConfig, func() {
		if err := repo.ClosePostgresConnection(); err != nil {
			slog.Error("app: can't close postgres connection", "err", err.Error())
		}
	})

	return &App{
		router:     router,
		repository: repo,
	}
}

func (a *App) Run() {
	a.router.Run()
}
