package router

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/medods/config"
	"github.com/avran02/medods/internal/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	c controller.Controller
	*chi.Mux
	config config.ServerConfig

	shutdownFunc func()
}

func getRoutes(c controller.Controller, apiPrefix string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	r.Get(apiPrefix+"/get-tokens", c.GetTokens)
	r.Post(apiPrefix+"/refresh-tokens", c.RefreshTokens)

	return r
}

func (r *Router) Run() {
	endpoint := r.config.Host + ":" + r.config.Port
	slog.Info("Server running on http://" + endpoint)
	s := http.Server{
		Addr:    endpoint,
		Handler: r.Mux,
	}

	s.RegisterOnShutdown(r.shutdownFunc)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("can't start server:\n", err)
	}
}

func New(c controller.Controller, conf config.ServerConfig, shutdownFunc func()) *Router {
	return &Router{
		c:      c,
		Mux:    getRoutes(c, conf.APIPrefix),
		config: conf,

		shutdownFunc: shutdownFunc,
	}
}
