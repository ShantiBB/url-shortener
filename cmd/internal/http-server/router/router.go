package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"url-shortener/cmd/internal/http-server/handlers/url/save"
	"url-shortener/cmd/internal/service"

	mwLogger "url-shortener/cmd/internal/http-server/middleware/logger"
)

func URLRouter(router *chi.Mux, log *slog.Logger, service *service.URLService) {
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, service))
}
